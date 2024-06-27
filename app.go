package main

import (
	"context"
	"os"
	"robtec/webapp/router"
)

var defaultPort = "8080"

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = defaultPort
	}

	ctx := context.Background()

	router.RunHTTPServer(ctx, port)
}
