package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Username  string             `json:"username" bson:"username"`
	Password  string             `json:"password" bson:"password"`
	Followers []string           `json:"followers" bson:"followers"`
	Following []string           `json:"following" bson:"following"`
	Role      string             `json:"role" bson:"role"`
}

type UserRepository interface {
	CreateUser(User) error
	FindByUsername(string) (User, error)
	FindUserByID(primitive.ObjectID) (User, error)
	FindAllUsers() ([]User, error)
	UpdateUser(User) error
	DeleteUser(primitive.ObjectID) error
	FollowUser(primitive.ObjectID, primitive.ObjectID) error
}
