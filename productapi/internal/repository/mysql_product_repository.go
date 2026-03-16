package repository

import (
	"api-estudo/internal/entities"

	"gorm.io/gorm"
)

type MySQLProductRepository struct {
	db *gorm.DB
}

func NewMySQLProductRepository(db *gorm.DB) *MySQLProductRepository {
	return &MySQLProductRepository{db: db}
}

func (r *MySQLProductRepository) GetAll() []entities.Product {
	var products []entities.Product

	if err := r.db.Find(&products).Error; err != nil {
		return []entities.Product{}
	}
	return products
}

func (r *MySQLProductRepository) GetByID(id string) (entities.Product, bool) {
	var p entities.Product

	if err := r.db.First(&p, "id = ?", id).Error; err != nil {
		return entities.Product{}, false
	}
	return p, true
}

func (r *MySQLProductRepository) Create(product entities.Product) error {

	return r.db.Create(&product).Error
}

func (r *MySQLProductRepository) Update(id string, product entities.Product) error {

	return r.db.Model(&entities.Product{}).
		Where("id = ?", id).
		Updates(product).Error
}

func (r *MySQLProductRepository) Delete(id string) error {

	return r.db.Delete(&entities.Product{}, "id = ?", id).Error
}
