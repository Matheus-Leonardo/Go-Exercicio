package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"

	"api-estudo/internal/config"
	"api-estudo/internal/entities"
	"api-estudo/internal/helpers"
	"api-estudo/internal/repository"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Falha ao carregar configurações: %v", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Erro ao abrir banco de dados: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Banco de dados inacessível: %v", err)
	}

	repo := repository.NewMySQLProductRepository(db)

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
		var p entities.Product
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			helpers.WriteJsonResponse(w, http.StatusBadRequest, map[string]string{"error": "JSON inválido"})
			return
		}

		if err := repo.Create(p); err != nil {
			helpers.WriteJsonResponse(w, http.StatusInternalServerError, map[string]string{"error": "Erro ao criar produto"})
			return
		}

		helpers.WriteJsonResponse(w, http.StatusCreated, p)
	})

	r.Put("/products/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		var p entities.Product
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			helpers.WriteJsonResponse(w, http.StatusBadRequest, map[string]string{"error": "JSON inválido"})
			return
		}

		if err := repo.Update(id, p); err != nil {
			helpers.WriteJsonResponse(w, http.StatusInternalServerError, map[string]string{"error": "Erro ao atualizar produto"})
			return
		}

		helpers.WriteJsonResponse(w, http.StatusOK, p)
	})

	r.Delete("/products/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if _, ok := repo.GetByID(id); !ok {
			helpers.WriteJsonResponse(w, http.StatusNotFound, map[string]string{"error": "produto não encontrado"})
			return
		}

		if err := repo.Delete(id); err != nil {
			helpers.WriteJsonResponse(w, http.StatusInternalServerError, map[string]string{"error": "Erro ao deletar produto"})
			return
		}

		w.WriteHeader(http.StatusNoContent)
	})

	fmt.Printf("API rodando em ambiente [%s] na porta 8081...\n", cfg.Database.Host)
	http.ListenAndServe(":8081", r)
}
