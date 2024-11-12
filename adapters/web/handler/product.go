package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Soter-Tec/go-hexagonal/adapters/dto"
	"github.com/Soter-Tec/go-hexagonal/aplication"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func MakeProductHandler(r *mux.Router, n *negroni.Negroni, service aplication.ProductServiceInterface) {
	r.Handle("/product/{id}", n.With(
		negroni.Wrap((getProduct(service))),
	)).Methods("GET", "OPTIONS")

	r.Handle("/product", n.With(
		negroni.Wrap((createProduct(service))),
	)).Methods("POST", "OPTIONS")

	r.Handle("/product/{id}", n.With(
		negroni.Wrap((enableDisableProduct(service))),
	)).Methods("PUT", "OPTIONS")
}

func getProduct(service aplication.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]
		product, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func createProduct(service aplication.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var productDto dto.Product
		err := json.NewDecoder(r.Body).Decode(&productDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
		product, err := service.Create(productDto.Name, productDto.Price)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

	})
}

func enableDisableProduct(service aplication.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]
		product, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write(jsonError(err.Error()))
			return
		}
		if product.GetStatus() == aplication.DISABLED {
			product, err := service.Enable(product)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(jsonError(err.Error()))
				return
			}
			err = json.NewEncoder(w).Encode(product)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(jsonError(err.Error()))
				return
			}
		} else {

			product, err := service.Disable(product)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(jsonError(err.Error()))
				return
			}
			err = json.NewEncoder(w).Encode(product)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(jsonError(err.Error()))
				return
			}
		}
		// var productDto dto.Product
		// err := json.NewDecoder(r.Body).Decode(&productDto)
		// if err != nil {
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	w.Write(jsonError(err.Error()))
		// 	return
		// }
		// product, err := service.Create(productDto.Name, productDto.Price)
		// if err != nil {
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	w.Write(jsonError(err.Error()))
		// 	return
		// }
		// err = json.NewEncoder(w).Encode(product)
		// if err != nil {
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	w.Write(jsonError(err.Error()))
		// 	return
		// }

	})
}
