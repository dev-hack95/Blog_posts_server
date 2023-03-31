package main

import (
	"log"
	"net/http"

	"github.com/dev-hack95/Blog_posts_server/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.CreateNewPost(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", r))
}
