package aplication_test

import (
	"testing"

	"github.com/Soter-Tec/go-hexagonal/aplication"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enabled(t *testing.T) {
	product := aplication.Product{}
	product.Name = "Hello"
	product.Status = aplication.DISABLED
	product.Price = 10

	err := product.Enable()
	if err != nil {
		require.Nil(t, err)
	}

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "The price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := aplication.Product{}
	product.Name = "Hello"
	product.Status = aplication.ENABLED
	product.Price = 10

	err := product.Disable()
	if err != nil {
		require.Nil(t, err)
	}

	product.Price = 0
	err = product.Disable()
	require.Equal(t, "The price must be greater than zero to disable the product", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := aplication.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Hello"
	product.Price = 10
	product.Status = aplication.DISABLED

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "The status must be enabled or disabled", err.Error())

	product.Status = aplication.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "The price must be greater or equal zero", err.Error())

}
