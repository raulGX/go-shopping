package cart

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/unrolled/render"
)

var (
	formatter = render.New(render.Options{
		IndentJSON: true,
	})
)

var productRepo = NewInMemoryProductRepository()

func TestCreateProduct(t *testing.T) {
	client := &http.Client{}
	server := httptest.NewServer(
		http.HandlerFunc(createProductHandler(formatter, productRepo)))
	defer server.Close()

	body := []byte("{\"name\": \"avocado\"}")
	body2 := []byte("{\"empty\": \"inside\"}")

	req, err := http.NewRequest("POST", server.URL, bytes.NewBuffer(body))
	if err != nil {
		t.Errorf("Error in creating POST request for createProductHandler: %v", err)
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		t.Errorf("Error in POST to createProductHandler: %v", err)
	}

	payload, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error in reading response body: %v", err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Errorf("Expected response status 201, received %s", res.Status)
	}

	req, err = http.NewRequest("POST", server.URL, bytes.NewBuffer(body2))
	if err != nil {
		t.Errorf("Error in creating POST request for createProductHandler: %v", err)
	}

	res, err = client.Do(req)

	payload, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error in reading response body: %v", err)
	}

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected response status 400, received %s", res.Status)
	}

	fmt.Printf("Payload: %s", string(payload))
}

// TestListProducts will only work if ran with the other tests
// It expects insertions to be made
// TODO clear the list and then make this into an integration test
func TestListProducts(t *testing.T) {
	client := &http.Client{}
	server := httptest.NewServer(
		http.HandlerFunc(listProductsHandler(formatter, productRepo)))
	defer server.Close()

	req, err := http.NewRequest("GET", server.URL, bytes.NewBuffer([]byte{}))
	if err != nil {
		t.Errorf("Error in creating POST request for createProductHandler: %v", err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		t.Errorf("Error in POST to createProductHandler: %v", err)
	}

	payload, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error in reading response body: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected response status 200, received %s", res.Status)
	}

	results := &Products{}
	err = json.Unmarshal(payload, &results)
	if err != nil {
		t.Error("Payload is not valid json")
	}

	if len(*results) != 1 {
		// not sure about this
		// results from the test above (addproduct)
		t.Error("Number of results doesn't match the required value")
	}
	fmt.Printf("Payload: %s", string(payload))
}
