package repository

import (
	"api-estudo/internal/entities"
	"database/sql"
	"log"
)

type MySQLProductRepository struct {
	db *sql.DB
}

// Construtor: recebe a conexão *sql.DB aberta no main.go
func NewMySQLProductRepository(db *sql.DB) *MySQLProductRepository {
	return &MySQLProductRepository{db: db}
}

// Implementações dos métodos da interface ProductRepository

func (r *MySQLProductRepository) GetAll() []entities.Product {
	rows, err := r.db.Query("SELECT id, name, type, quantity FROM product")
	if err != nil {
		log.Println("Erro ao buscar produtos:", err)
		return []entities.Product{}
	}
	defer rows.Close()

	var products []entities.Product
	for rows.Next() {
		var p entities.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Type, &p.Quantity); err == nil {
			products = append(products, p)
		}
	}
	return products
}

func (r *MySQLProductRepository) GetByID(id string) (entities.Product, bool) {
	var p entities.Product
	err := r.db.QueryRow("SELECT id, name, type, quantity FROM product WHERE id = ?", id).
		Scan(&p.ID, &p.Name, &p.Type, &p.Quantity)
	if err != nil {
		return entities.Product{}, false
	}
	return p, true
}

func (r *MySQLProductRepository) Create(product entities.Product) error {
	_, _ = r.db.Exec("INSERT INTO product (id, name, type, quantity) VALUES (?, ?, ?, ?)",
		product.ID, product.Name, product.Type, product.Quantity)
	return nil
}

func (r *MySQLProductRepository) Update(id string, product entities.Product) error {
	_, _ = r.db.Exec("UPDATE product SET name=?, type=?, quantity=? WHERE id=?",
		product.Name, product.Type, product.Quantity, id)
	return nil
}

func (r *MySQLProductRepository) Delete(id string) error {
	_, _ = r.db.Exec("DELETE FROM product WHERE id=?", id)
	return nil
}
