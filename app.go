package main

import (
	"context"
	"flag"
	"robtec/webapp/router"
)

func main() {

	var (
		port string
	)

	flag.StringVar(&port, "port", "8080", "http port")

	flag.Parse()

	ctx := context.Background()

	router.RunHTTPServer(ctx, port)
}
