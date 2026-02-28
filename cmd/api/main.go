package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	// Importando seus próprios pacotes internos
	"api-estudo/internal/entities"
	"api-estudo/internal/helpers"
)

// Simulação de banco de dados em memória
var products = map[string]entities.Product{
	"1": {ID: "1", Name: "Guitarra Fender", Type: "Instrumento", Quantity: 5},
}

func main() {
	r := chi.NewRouter()

	// Middlewares para monitorar a API no terminal
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Rota para listagem de produtos (GET)
	r.Get("/products", func(w http.ResponseWriter, r *http.Request) {
		helpers.WriteJsonResponse(w, http.StatusOK, products)
	})

	fmt.Println("API de Estoque Reiniciada na porta 8081...")
	http.ListenAndServe(":8081", r)
}
