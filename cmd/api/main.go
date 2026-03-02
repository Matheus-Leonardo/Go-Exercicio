package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"api-estudo/internal/entities"
	"api-estudo/internal/helpers"
)

var products = map[string]entities.Product{
	"1": {ID: "1", Name: "Guitarra Fender", Type: "Instrumento", Quantity: 5},
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/products", func(w http.ResponseWriter, r *http.Request) {
		helpers.WriteJsonResponse(w, http.StatusOK, products)
	})

	r.Get("/products/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id") // Extrai o ID da URL usando o recurso do Chi
		product, ok := products[id] // O "idioma" do Go para verificar se a chave existe
		if !ok {
			helpers.WriteJsonResponse(w, http.StatusNotFound, map[string]string{"error": "produto não encontrado"})
			return
		}
		helpers.WriteJsonResponse(w, http.StatusOK, product)
	})

	r.Post("/products", func(w http.ResponseWriter, r *http.Request) {
		var newProduct entities.Product

		err := json.NewDecoder(r.Body).Decode(&newProduct)
		if err != nil {
			helpers.WriteJsonResponse(w, http.StatusBadRequest, map[string]string{"error": "JSON inválido"})
			return
		}

		products[newProduct.ID] = newProduct

		helpers.WriteJsonResponse(w, http.StatusCreated, newProduct)
	})

	fmt.Println("API Completa (Ep. 1) rodando na porta 8081...")
	http.ListenAndServe(":8081", r)
}
