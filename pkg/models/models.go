package models

import (
	"fmt"

	"github.com/dev-hack95/Blog_posts_server/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Post struct {
	gorm.Model
	ID        uint64 `gorm:"primary_key;AUTO_INCREMENT;NOT NULL" json:"id"`
	Title     string `gorm:"NOT NULL" json:"title"`
	Content   string `gorm:"NOT NULL" json:"context" `
	Published bool   `gorm:"default:true" json:"published"`
	Author    Author `gorm:"embedded;NOT NULL" json:"author"`
}

type Author struct {
	Firstname string `gorm:"NOT NULL" json:"firstname"`
	Lastname  string `gorm:"NOT NULL" json:"lastname"`
}

func init() {
	// Connection to database
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Post{})
}

func (b *Post) CreatePosts() *Post {
	fmt.Println(&b)
	db.Create(&b)
	return b
}

func GetAllPosts() []Post {
	var posts []Post
	db.Find(&posts)
	return posts
}

func GetPostById(ID int64) (*Post, *gorm.DB) {
	var getPost Post
	db := db.Where("ID=?", ID).Find(&getPost)
	return &getPost, db
}

func DeletePosts(ID int64) *Post {
	var deletePost Post
	db.Where("ID=?", ID).Unscoped().Delete(&deletePost)
	return &deletePost
}
