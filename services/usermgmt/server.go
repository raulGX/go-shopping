package usermgmt

import (
	sql "database/sql"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func AddRoutes(mx *mux.Router, db *sql.DB) {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	repository := NewUserPostgresRepository(db)
	initRoutes(mx, formatter, repository)
}

func initRoutes(mx *mux.Router, formatter *render.Render, r UserRepository) {
	mx.HandleFunc("/users/{username}", getUserHandler(formatter, r)).Methods("GET")
	mx.HandleFunc("/users", createUserHandler(formatter, r)).Methods("POST")
}
