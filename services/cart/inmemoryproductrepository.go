package cart

var products = make(Products, 0)

type InMemoryProductRepository struct {
}

func NewInMemoryProductRepository() *InMemoryProductRepository {
	return &InMemoryProductRepository{}
}

func (r *InMemoryProductRepository) GetProducts() (Products, error) {
	return products, nil
}

func (r *InMemoryProductRepository) AddProduct(p Product) error {
	products = append(products, p)
	return nil
}
