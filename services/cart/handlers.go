package cart

import (
	"net/http"

	"github.com/unrolled/render"
)

func createProductHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"this is a test"})
	}
}
