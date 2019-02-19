package usermgmt

import (
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func AddRoutes(mx *mux.Router) {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	initRoutes(mx, formatter)
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/products", createUserHandler(formatter)).Methods("POST")
}
