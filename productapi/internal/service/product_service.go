package service

import (
	"api-estudo/internal/entities"
	"api-estudo/internal/repository"
	"errors"
	"strings"
)

type ProductService interface {
	Create(p entities.Product) error
	GetAll() []entities.Product
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) Create(p entities.Product) error {

	if strings.TrimSpace(p.Name) == "" {
		return errors.New("O nome do produto é obrigatório")
	}
	return s.repo.Create(p)
}

func (s *productService) GetAll() []entities.Product {
	return s.repo.GetAll()
}
