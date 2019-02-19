package cart

type AddProductRequest struct {
	Name string `json:"name" bson:"name"`
}

func (a AddProductRequest) IsValid() bool {
	valid := true
	if a.Name == "" {
		valid = false
	}
	return valid
}

type DBConfig struct {
	IP     string
	DBName string
}

type ProductRepository interface {
	AddProduct(Product) error
	GetProducts() (Products, error)
}
