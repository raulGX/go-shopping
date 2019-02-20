package usermgmt

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/unrolled/render"
)

func createUserHandler(formatter *render.Render, r UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		payload, err := ioutil.ReadAll(req.Body)
		if err != nil {
			formatter.Text(w, http.StatusUnprocessableEntity, "Failed to parse request body")
			return
		}

		var newUserReq UserCreateRequest
		err = json.Unmarshal(payload, &newUserReq)
		if err != nil {
			formatter.Text(w, http.StatusUnprocessableEntity, "Failed to parse request body")
			return
		}

		if !newUserReq.IsValid() {
			formatter.Text(w, http.StatusBadRequest, "Payload is not valid") // TODO should treat error better
			return
		}

		err = r.AddUser(newUserReq)
		if err != nil {
			formatter.Text(w, http.StatusBadRequest, "User with the same username already exists")
			return
		}

		formatter.JSON(w, 201, struct{}{})
	}
}

func getUserHandler(formatter *render.Render, r UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// mx.HandleFunc("/matches/{id}", getMatchDetailsHandler(formatter, repo)).Methods("GET")
		vars := mux.Vars(req)
		username := vars["username"]
		user, err := r.GetUserByUsername(username)
		fmt.Printf("%v", username)
		if err != nil {
			formatter.JSON(w, 404, struct{ Err string }{"Could not find the user"})
			return
		}
		formatter.JSON(w, 200, user)
	}
}
