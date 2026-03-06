package repository

import "api-estudo/internal/entities"

// ProductRepository define o contrato (interface) do que um repositório de produtos deve fazer.
type ProductRepository interface {
	GetAll() map[string]entities.Product
	GetByID(id string) (entities.Product, bool)
	Create(product entities.Product)
	Update(id string, product entities.Product)
	Delete(id string)
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

// Implementações dos métodos (A lógica que estava na main vem para cá)
func (r *MapProductRepository) GetAll() map[string]entities.Product {
	return r.products
}

func (r *MapProductRepository) GetByID(id string) (entities.Product, bool) {
	product, ok := r.products[id]
	return product, ok
}

func (r *MapProductRepository) Create(product entities.Product) {
	r.products[product.ID] = product
}

func (r *MapProductRepository) Update(id string, product entities.Product) {
	r.products[id] = product
}

func (r *MapProductRepository) Delete(id string) {
	delete(r.products, id)
}
