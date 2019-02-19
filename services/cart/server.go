package cart

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	mgo "gopkg.in/mgo.v2"
)

// Maybe mock for in memory
const SERVER = "mongodb://localhost:27017"

// DBNAME the name of the DB instance
const DBNAME = "store"

var dbcfg = &DBConfig{SERVER, DBNAME}

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{

		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	//create session here
	session, err := mgo.Dial(dbcfg.IP)
	if err != nil {
		panic("Could not connect to db")
	}
	productRepo := NewMongoProductRepository(session, dbcfg)
	initRoutes(mx, formatter, productRepo)
	n.UseHandler(mx)

	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render, r ProductRepository) {
	mx.HandleFunc("/products", listProductsHandler(formatter, r)).Methods("GET")
	mx.HandleFunc("/products", createProductHandler(formatter, r)).Methods("POST")
}
