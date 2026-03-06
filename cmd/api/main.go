package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"api-estudo/internal/entities"
	"api-estudo/internal/helpers"
	"api-estudo/internal/repository"
)

func main() {
	repo := repository.NewMapProductRepository()

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/products", func(w http.ResponseWriter, r *http.Request) {
		helpers.WriteJsonResponse(w, http.StatusOK, repo.GetAll())
	})

	r.Get("/products/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		product, ok := repo.GetByID(id)
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

		repo.Create(newProduct)

		helpers.WriteJsonResponse(w, http.StatusCreated, newProduct)
	})

	r.Put("/products/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if _, ok := repo.GetByID(id); !ok {
			helpers.WriteJsonResponse(w, http.StatusNotFound, map[string]string{"error": "produto inexistente"})
			return
		}

		var updatedProduct entities.Product
		if err := json.NewDecoder(r.Body).Decode(&updatedProduct); err != nil {
			fmt.Println("Erro ao decodificar JSON do PUT:", err)
			helpers.WriteJsonResponse(w, http.StatusBadRequest, map[string]string{"error": "JSON inválido"})
			return
		}

		updatedProduct.ID = id
		repo.Update(id, updatedProduct)

		fmt.Printf("Produto %s atualizado com sucesso!\n", id)
		helpers.WriteJsonResponse(w, http.StatusOK, updatedProduct)
	})

	r.Delete("/products/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if _, ok := repo.GetByID(id); !ok {
			helpers.WriteJsonResponse(w, http.StatusNotFound, map[string]string{"error": "produto não encontrado"})
			return
		}
		delete(repo.GetAll(), id)
		w.WriteHeader(http.StatusNoContent)
	})

	fmt.Println("API Completa (Ep. 1) rodando na porta 8081...")
	http.ListenAndServe(":8081", r)
}
