package routes

import (
	"github.com/dev-hack95/Blog_posts_server/pkg/controllers"
	"github.com/gorilla/mux"
)

var CreateNewPost = func(router *mux.Router) {
	router.HandleFunc("/posts/", controllers.GetAllPosts).Methods("GET")
	router.HandleFunc("/posts/", controllers.CreatePosts).Methods("POST")
	router.HandleFunc("/posts/{id}", controllers.GetPostById).Methods("GET")
	router.HandleFunc("/posts/{id}", controllers.UpdatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", controllers.DeletePosts).Methods("DELETE")
}
