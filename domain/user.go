package domain

type User struct {
	Username string `json:"username" binding:"required" bson:"username"`
	Password string `json:"password" binding:"required" bson:"password"`
	Followers []User `json:"followers" bson:"followers"`
	Following []User `json:"following" bson:"following"`
	Role string `json:"role" bson:"role"`
}

type UserRepository interface {
	CreateUser(User) error
	FindByUsername(string) (User, error)
	FindAll() ([]User, error)
	UpdateUser(User) error
	DeleteUser(string) error
}

