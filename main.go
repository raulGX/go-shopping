package main

import (
	"os"

	cartservice "github.com/raulgx/go-shopping/services/cart"
)

func main() {
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "3000"
	}

	server := cartservice.NewServer()

	server.Run(":" + port)
}
