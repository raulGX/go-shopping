package usermgmt

import (
	"bytes"
	"encoding/json"
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

func Test_createUserHandler(t *testing.T) {
	client := &http.Client{}
	server := httptest.NewServer(http.HandlerFunc(createUserHandler(formatter)))
	defer server.Close()

	body := []byte(`{"username": "macncheese@food.com", "password": "longenoughpassword"}`)
	badRequests := [][]byte{
		[]byte(`{"username": "", "password": "longenoughpassword"}`),
		[]byte(`{"username": "peter@pan.com", "password": "smol"}`),
	}

	req, err := http.NewRequest("POST", server.URL, bytes.NewBuffer(body))
	if err != nil {
		t.Errorf("Error in creating POST request for createUserHandler: %v", err)
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		t.Errorf("Error in receiving response for createUserHandler: %v", err)
	}

	payload, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error in processing response for createUserHandler: %v", err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Errorf("Expecting status 201, got: %v", res.StatusCode)
	}

	var createUserRequest UserCreateRequest
	err = json.Unmarshal(payload, &createUserRequest)
	if err != nil {
		t.Errorf("Error in processing payload for createUserHandler: %v", err)
	}

	// bad requests
	for _, test := range badRequests {
		req, err := http.NewRequest("POST", server.URL, bytes.NewBuffer(test))
		if err != nil {
			t.Errorf("Error in creating POST request for createUserHandler: %v", err)
		}

		req.Header.Add("Content-Type", "application/json")

		res, err := client.Do(req)
		if err != nil {
			t.Errorf("Error in receiving response for createUserHandler: %v", err)
		}

		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("Expecting status 400, got: %v", res.StatusCode)
		}
	}
}
