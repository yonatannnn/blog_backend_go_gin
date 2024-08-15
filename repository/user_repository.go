package repository

import (
	"blog_api/domain"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	count, err := ur.databaseCollection.CountDocuments(ur.context, bson.D{})
	if err != nil {
		return err
	}

	if count == 0 {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}
	
	if _ , err := ur.FindByUsername(user.Username) ; err == nil {
		return errors.New("User already exists")
	}
	_ , err = ur.databaseCollection.InsertOne(ur.context, user)
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

func (ur *userRepository) FindAllUsers() ([]domain.User, error) {
	var users []domain.User
	cursor, err := ur.databaseCollection.Find(ur.context, bson.D{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(ur.context) {
		var user domain.User
		cursor.Decode(&user)
		users = append(users, user)
	}

	return users, nil
}

func (ur *userRepository) UpdateUser(user domain.User) error {
	// Implement logic to update a user in the database
	return nil
}

func (ur *userRepository) DeleteUser(objID primitive.ObjectID) error {
	// Implement logic to delete a user from the database
	filter := bson.D{{"_id", objID}}
	err := ur.databaseCollection.FindOneAndDelete(ur.context , filter)
	if err.Err() != nil {
		return errors.New("Failed to delete user")
	}
	return nil
}


