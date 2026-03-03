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

	r.Put("/products/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		// Verificamos se o item existe antes de tentar mudar
		if _, ok := products[id]; !ok {
			helpers.WriteJsonResponse(w, http.StatusNotFound, map[string]string{"error": "produto inexistente"})
			return
		}

		var updatedProduct entities.Product
		if err := json.NewDecoder(r.Body).Decode(&updatedProduct); err != nil {
			fmt.Println("Erro ao decodificar JSON do PUT:", err) // Log para você ver no terminal
			helpers.WriteJsonResponse(w, http.StatusBadRequest, map[string]string{"error": "JSON inválido"})
			return
		}

		updatedProduct.ID = id        // Forçamos o ID da URL para garantir consistência
		products[id] = updatedProduct // Atualizamos a "gaveta" correta do mapa

		fmt.Printf("Produto %s atualizado com sucesso!\n", id)
		helpers.WriteJsonResponse(w, http.StatusOK, updatedProduct)
	})

	r.Delete("/products/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if _, ok := products[id]; !ok {
			helpers.WriteJsonResponse(w, http.StatusNotFound, map[string]string{"error": "produto não encontrado"})
			return
		}
		delete(products, id)                // Função nativa do Go para remover de mapas
		w.WriteHeader(http.StatusNoContent) // 204: Sucesso, mas sem corpo de resposta
	})

	fmt.Println("API Completa (Ep. 1) rodando na porta 8081...")
	http.ListenAndServe(":8081", r)
}
