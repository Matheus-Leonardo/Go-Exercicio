package service

import (
	"api-estudo/internal/entities"
	"api-estudo/internal/repository"
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
)

type ProductService interface {
	Create(ctx context.Context, p entities.Product) error
	GetAll(ctx context.Context) []entities.Product
	Search(ctx context.Context, productType string) []entities.Product
	GetByID(ctx context.Context, id string) (entities.Product, error)
	Update(ctx context.Context, id string, p entities.Product) error
	Delete(ctx context.Context, id string) error
}

type productService struct {
	repo repository.ProductRepository
}

// NewProductService injeta o repositório no serviço
func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) Create(ctx context.Context, p entities.Product) error {
	if strings.TrimSpace(p.Name) == "" {
		return errors.New("O nome do produto é obrigatório")
	}

	id, err := uuid.NewUUID()
	if err != nil {
		return errors.New("erro ao gerar ID")
	}
	p.ID = id.String()
	return s.repo.Create(ctx, p)
}

func (s *productService) GetAll(ctx context.Context) []entities.Product {
	return s.repo.GetAll(ctx)
}
func (s *productService) GetByID(ctx context.Context, id string) (entities.Product, error) {
	product, ok := s.repo.GetByID(ctx, id)
	if !ok {
		return entities.Product{}, errors.New("produto não encontrado")
	}
	return product, nil
}

func (s *productService) Search(ctx context.Context, productType string) []entities.Product {
	return s.repo.Search(ctx, productType)
}

func (s *productService) Update(ctx context.Context, id string, p entities.Product) error {
	if strings.TrimSpace(p.Name) == "" {
		return errors.New("O nome do produto é obrigatório")
	}
	return s.repo.Update(ctx, id, p)
}

func (s *productService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
