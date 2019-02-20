package cart

import (
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	mgo "gopkg.in/mgo.v2"
)

const SERVER = "mongodb://localhost:27017"

// DBNAME the name of the DB instance
const DBNAME = "store"

var dbcfg = &DBConfig{SERVER, DBNAME}

// AddRoutes adds routes to existing router mx.
func AddRoutes(mx *mux.Router) {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	//create session here
	session, err := mgo.Dial(dbcfg.IP)
	// should close the session on exit?
	if err != nil {
		panic("Could not connect to db")
	}
	productRepo := NewMongoProductRepository(session, dbcfg)

	initRoutes(mx, formatter, productRepo)
}

func initRoutes(mx *mux.Router, formatter *render.Render, r ProductRepository) {
	mx.HandleFunc("/products", listProductsHandler(formatter, r)).Methods("GET")
	mx.HandleFunc("/products", createProductHandler(formatter, r)).Methods("POST")
}
