package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Username  string   `json:"username" binding:"required" bson:"username"`
	Password  string   `json:"password" binding:"required" bson:"password"`
	Followers []string `json:"followers" bson:"followers"`
	Following []string `json:"following" bson:"following"`
	Role      string   `json:"role" bson:"role"`
}

type UserRepository interface {
	CreateUser(User) error
	FindByUsername(string) (User, error)
	FindAllUsers() ([]User, error)
	UpdateUser(User) error
	DeleteUser(primitive.ObjectID) error
}
