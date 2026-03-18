package service

import (
	"api-estudo/internal/entities"
	"api-estudo/internal/repository"
	"errors"
	"strings"
)

// ProductService define o que o negócio espera de um produto
type ProductService interface {
	Create(p entities.Product) error
	GetAll() []entities.Product
	GetByID(id string) (entities.Product, error) // retorna product + error
	Update(id string, p entities.Product) error
	Delete(id string) error
}

type productService struct {
	// Cross-check: Depende da interface do repositório, não da implementação MySQL
	repo repository.ProductRepository
}

// NewProductService injeta o repositório no serviço
func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) Create(p entities.Product) error {
	// Remove espaços em branco antes de validar
	if strings.TrimSpace(p.Name) == "" {
		return errors.New("O nome do produto é obrigatório")
	}
	return s.repo.Create(p)
}

func (s *productService) GetAll() []entities.Product {
	return s.repo.GetAll()
}
func (s *productService) GetByID(id string) (entities.Product, error) {
	product, ok := s.repo.GetByID(id)
	if !ok {
		return entities.Product{}, errors.New("produto não encontrado")
	}
	return product, nil
}

func (s *productService) Update(id string, p entities.Product) error {
	if strings.TrimSpace(p.Name) == "" {
		return errors.New("O nome do produto é obrigatório")
	}
	return s.repo.Update(id, p)
}

func (s *productService) Delete(id string) error {
	return s.repo.Delete(id)
}
