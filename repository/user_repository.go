package repository

import (
	"blog_api/domain"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	databaseCollection *mongo.Collection
	context			context.Context
}

func NewUserRepository(collection *mongo.Collection) domain.UserRepository {
	return &userRepository{
		databaseCollection: collection,
		context: context.TODO(),
	}
}


func (ur *userRepository) CreateUser(user domain.User) error {
	
	_, err := ur.databaseCollection.InsertOne(ur.context, user)
	if err != nil {
		return errors.New("Failed to create user")
	}
	return nil
}

func (ur *userRepository) FindByUsername(username string) (domain.User, error) {
	var user domain.User
	err := ur.databaseCollection.FindOne(ur.context , bson.D{{"username", username}}).Decode(&user)
	if err != nil {
		return domain.User{}, errors.New("User not found")
	}
	return user, nil

}

func (ur *userRepository) FindAll() ([]domain.User, error) {
	// Implement logic to find all users in the database
	return []domain.User{}, nil
}

func (ur *userRepository) UpdateUser(user domain.User) error {
	// Implement logic to update a user in the database
	return nil
}

func (ur *userRepository) DeleteUser(id string) error {
	// Implement logic to delete a user by ID from the database
	return nil
}


