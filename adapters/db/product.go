package db

import (
	"database/sql"

	"github.com/Soter-Tec/go-hexagonal/aplication"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (p *ProductDb) Get(id string) (aplication.ProductInterface, error) {

	var product aplication.Product

	stmt, err := p.db.Prepare("select id, name, price, status from products where id = ?")
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}
	return &product, nil

}

func (p *ProductDb) Save(product aplication.ProductInterface) (aplication.ProductInterface, error) {
	var rows int
	p.db.QueryRow("select count(id) from products where id = ?", product.GetId()).Scan(&rows)
	if rows == 0 {
		_, err := p.create(product)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := p.update(product)
		if err != nil {
			return nil, err
		}
	}
	return product, nil
}

func (p *ProductDb) create(product aplication.ProductInterface) (aplication.ProductInterface, error) {

	stmt, err := p.db.Prepare("insert into products(id, name, price, status) values(?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(product.GetId(), product.GetName(), product.GetPrice(), product.GetStatus())
	if err != nil {
		return nil, err
	}
	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductDb) update(product aplication.ProductInterface) (aplication.ProductInterface, error) {

	stmt, err := p.db.Prepare("update products set name = ?, price = ?, status = ? where id = ?")
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(product.GetName(), product.GetPrice(), product.GetStatus(), product.GetId())
	if err != nil {
		return nil, err
	}
	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	return product, nil
}
