package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Title     string             `json:"title" binding:"required" bson:"title"`
	Content   string             `json:"content" binding:"required" bson:"content"`
	Author    User               `json:"author" bson:"author"`
	Likes     []string           `json:"likes" bson:"likes"`
	CreatedAt time.Time          `json:"creater_at" bson:"creater_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type PostRepository interface {
	CreatePost(Post) error
	FindPostById(primitive.ObjectID) (Post, error)
	FindAllPosts() ([]Post, error)
	UpdatePost(Post) error
	DeletePost(primitive.ObjectID) error
	LikePost(primitive.ObjectID, User) error
	UnlikePost(primitive.ObjectID, User) error
}
