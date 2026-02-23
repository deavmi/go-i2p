package web;

import "github.com/go-i2p/router";
import "fmt";

type WebServer struct {
	r *router.Router
}

func NewWebServer(r *router.Router) {
	fmt.Printf("Creating web server...\n")
}
