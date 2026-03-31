package repository

import (
	"api-estudo/internal/entities"
	"context"

	"gorm.io/gorm"
)

type MySQLProductRepository struct {
	db *gorm.DB
}

func NewMySQLProductRepository(db *gorm.DB) *MySQLProductRepository {
	return &MySQLProductRepository{db: db}
}

func (r *MySQLProductRepository) GetAll(ctx context.Context) []entities.Product {
	var products []entities.Product

	if err := r.db.WithContext(ctx).Find(&products).Error; err != nil {
		return []entities.Product{}
	}
	return products
}

func (r *MySQLProductRepository) GetByID(ctx context.Context, id string) (entities.Product, bool) {
	var p entities.Product

	if err := r.db.WithContext(ctx).First(&p, "id = ?", id).Error; err != nil {
		return entities.Product{}, false
	}
	return p, true
}

func (r *MySQLProductRepository) Create(ctx context.Context, product entities.Product) error {

	return r.db.WithContext(ctx).Create(&product).Error
}

func (r *MySQLProductRepository) Update(ctx context.Context, id string, product entities.Product) error {

	return r.db.WithContext(ctx).Model(&entities.Product{}).
		Where("id = ?", id).
		Updates(product).Error
}

func (r *MySQLProductRepository) Delete(ctx context.Context, id string) error {

	return r.db.WithContext(ctx).Delete(&entities.Product{}, "id = ?", id).Error
}
