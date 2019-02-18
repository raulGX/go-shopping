package cart

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/unrolled/render"
)

func createProductHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		payload, _ := ioutil.ReadAll(req.Body)

		var addProductReq AddProductRequest
		err := json.Unmarshal(payload, &addProductReq)

		if err != nil {
			formatter.Text(w, http.StatusBadRequest, "Failed to parse match request")
			return
		}

		if addProductReq.IsValid() != true {
			formatter.Text(w, http.StatusBadRequest, "Invalid Payload")
			return
		}

		newProduct := Product{Name: addProductReq.Name}
		err = ProductRepository{}.AddProduct(newProduct) // TODO INJECT
		if err != nil {
			formatter.Text(w, http.StatusBadRequest, "Failed to save to db")
		}

		formatter.JSON(w, http.StatusCreated, newProduct)
	}
}

func listProductsHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		products, err := ProductRepository{}.GetProducts()
		if err != nil {
			formatter.Text(w, http.StatusInternalServerError, "Failed to fetch products")
			return
		}
		formatter.JSON(w, http.StatusOK, products)
	}
}
