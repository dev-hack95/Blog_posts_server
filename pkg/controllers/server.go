package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dev-hack95/Blog_posts_server/pkg/models"
	"github.com/dev-hack95/Blog_posts_server/pkg/utils"
	"github.com/gorilla/mux"
)

var NewPost models.Post

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	newPosts := models.GetAllPosts()
	res, _ := json.Marshal(newPosts)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreatePosts(w http.ResponseWriter, r *http.Request) {
	createPost := &models.Post{}
	utils.ParseBody(r, createPost)
	b := createPost.CreatePosts()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetPostById(w http.ResponseWriter, r *http.Request) {
	postId := mux.Vars(r)["id"]
	ID, err := strconv.ParseInt(postId, 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	postDetails, _ := models.GetPostById(ID)
	res, _ := json.Marshal(postDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeletePosts(w http.ResponseWriter, r *http.Request) {
	postId := mux.Vars(r)["id"]
	ID, err := strconv.ParseInt(postId, 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	postDetalils := models.DeletePosts(ID)

	res, _ := json.Marshal(postDetalils)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	var updatePost = &models.Post{}
	utils.ParseBody(r, updatePost)
	postId := mux.Vars(r)["id"]
	ID, err := strconv.ParseInt(postId, 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	postDetails, db := models.GetPostById(ID)

	if updatePost.Title != "" {
		postDetails.Title = updatePost.Title
	}

	if updatePost.Content != "" {
		postDetails.Content = updatePost.Content
	}

	if updatePost.Author.Firstname != "" {
		postDetails.Author.Firstname = updatePost.Author.Firstname
	}

	if updatePost.Author.Lastname != "" {
		postDetails.Author.Lastname = updatePost.Author.Lastname
	}

	db.Save(&postDetails)
	res, _ := json.Marshal(postDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
