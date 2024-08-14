package domain

import "time"

type Post struct {
	ID        string `json:"id" bson:"_id"`
	Content   string `json:"content" binding:"required" bson:"content"`
	Author    User  `json:"author" binding:"required" bson:"author"`
	Likes     []User `json:"likes" bson:"likes"`
	CreaterAt time.Time `json:"creater_at" bson:"creater_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}


type PostRepository interface {
	CreatePost(Post) error
	FindById(string) (Post, error)
	FindAll() ([]Post, error)
	UpdatePost(Post) error
	DeletePost(string) error
}