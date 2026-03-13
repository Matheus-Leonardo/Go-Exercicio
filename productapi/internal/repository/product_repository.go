package repository

import "api-estudo/internal/entities"

// ProductRepository define o contrato (interface) do que um repositório de produtos deve fazer.
type ProductRepository interface {
	GetAll() []entities.Product
	GetByID(id string) (entities.Product, bool)
	Create(product entities.Product) error
	Update(id string, product entities.Product) error
	Delete(id string) error
}

// MapProductRepository é a implementação real usando o Mapa na RAM (Mock).
type MapProductRepository struct {
	products map[string]entities.Product
}

// NewMapProductRepository é o construtor que limpa e inicializa nosso "banco" em memória.
func NewMapProductRepository() *MapProductRepository {
	return &MapProductRepository{
		products: map[string]entities.Product{
			"1": {ID: "1", Name: "Guitarra Fender", Type: "Instrumento", Quantity: 5},
		},
	}
}

func (r *MapProductRepository) GetAll() []entities.Product {
	products := []entities.Product{}
	for _, p := range r.products {
		products = append(products, p)
	}
	return products
}

func (r *MapProductRepository) GetByID(id string) (entities.Product, bool) {
	product, ok := r.products[id]
	return product, ok
}

func (r *MapProductRepository) Create(product entities.Product) error {
	r.products[product.ID] = product
	return nil
}

func (r *MapProductRepository) Update(id string, product entities.Product) error {
	r.products[id] = product
	return nil
}

func (r *MapProductRepository) Delete(id string) error {
	delete(r.products, id)
	return nil
}
