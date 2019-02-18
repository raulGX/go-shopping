package cart

import (
	"bytes"
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

func TestCreateProduct(t *testing.T) {
	client := &http.Client{}
	server := httptest.NewServer(
		http.HandlerFunc(createProductHandler(formatter)))
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

func TestListProducts(t *testing.T) {
	client := &http.Client{}
	server := httptest.NewServer(
		http.HandlerFunc(listProductsHandler(formatter)))
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

	fmt.Printf("Payload: %s", string(payload))
}
