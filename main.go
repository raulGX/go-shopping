package main

import (
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	cartservice "github.com/raulGX/go-shopping/services/cart"
	usermgmtservice "github.com/raulGX/go-shopping/services/usermgmt"
)

func main() {
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "3000"
	}

	n := negroni.Classic()
	mx := mux.NewRouter()
	db := usermgmtservice.NewPostgresConnection()
	defer db.Close()
	cartservice.AddRoutes(mx)
	usermgmtservice.AddRoutes(mx, db)

	n.UseHandler(mx)

	n.Run(":" + port)
}
