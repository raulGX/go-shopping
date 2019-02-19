package usermgmt

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/unrolled/render"
)

func createUserHandler(formatter *render.Render) http.HandlerFunc {
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

		// TODO add repo.addUser

		formatter.JSON(w, 201, struct{ Text string }{"oke"})
	}
}
