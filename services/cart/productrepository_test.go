package cart

import (
	"reflect"
	"testing"
)

func TestProductRepository_GetProducts(t *testing.T) {
	tests := []struct {
		name    string
		r       ProductRepository
		want    Products
		wantErr bool
	}{
		// TODO setup CICD env
		// {"empty", ProductRepository{}, Products{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ProductRepository{}
			got, err := r.GetProducts()
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductRepository.GetProducts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductRepository.GetProducts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductRepository_AddProduct(t *testing.T) {
	type args struct {
		p Product
	}
	tests := []struct {
		name    string
		r       ProductRepository
		args    args
		wantErr bool
	}{
		// TODO setup CICD env
		// {"product", ProductRepository{}, args{p: Product{Name: "avocado1"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ProductRepository{}
			if err := r.AddProduct(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("ProductRepository.AddProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
