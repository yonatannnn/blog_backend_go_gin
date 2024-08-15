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
	context            context.Context
}

func NewUserRepository(collection *mongo.Collection) domain.UserRepository {
	return &userRepository{
		databaseCollection: collection,
		context:            context.TODO(),
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

	if _, err := ur.FindByUsername(user.Username); err == nil {
		return errors.New("User already exists")
	}
	_, err = ur.databaseCollection.InsertOne(ur.context, user)
	if err != nil {
		return errors.New("Failed to create user")
	}
	return nil
}

func (ur *userRepository) FindByUsername(username string) (domain.User, error) {
	var user domain.User
	err := ur.databaseCollection.FindOne(ur.context, bson.D{{"username", username}}).Decode(&user)
	if err != nil {
		return domain.User{}, errors.New("User not found")
	}
	return user, nil

}

func (ur *userRepository) FindUserByID(id primitive.ObjectID) (domain.User, error) {
	var user domain.User
	err := ur.databaseCollection.FindOne(ur.context, bson.D{{"_id", id}}).Decode(&user)
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
	filter := bson.D{bson.E{"_id", user.ID}}
	updatedFields := bson.D{}
	if user.Username != "" {
		updatedFields = append(updatedFields, bson.E{Key: "username", Value: user.Username})
	}
	if user.Password != "" {
		updatedFields = append(updatedFields, bson.E{Key: "password", Value: user.Password})
	}

	update := bson.D{
		bson.E{
			Key: "$set", Value: updatedFields},
	}

	_, err := ur.databaseCollection.UpdateOne(ur.context, filter, update)
	if err != nil {
		return errors.New("failed to update user!")
	}

	return nil
}

func (ur *userRepository) DeleteUser(objID primitive.ObjectID) error {
	// Implement logic to delete a user from the database
	filter := bson.D{{"_id", objID}}
	err := ur.databaseCollection.FindOneAndDelete(ur.context, filter)
	if err.Err() != nil {
		return errors.New("Failed to delete user")
	}
	return nil
}

func (ur *userRepository) FollowUser(followerId primitive.ObjectID, followeeID primitive.ObjectID) error {
	followerFilter := bson.D{bson.E{"_id", followerId}}
	followeeFilter := bson.D{bson.E{"_id", followeeID}}

	update := bson.D{
		bson.E{
			Key: "$push", Value: bson.E{Key: "followers", Value: followerId}},
	}

	update2 := bson.D{
		bson.E{
			Key: "$push", Value: bson.E{Key: "following", Value: followeeID}},
	}

	_, err1 := ur.databaseCollection.UpdateOne(ur.context, followerFilter, update2)
	if err1 != nil {
		return errors.New("failed in updating follower")
	}

	_, err2 := ur.databaseCollection.UpdateOne(ur.context, followeeFilter, update)
	if err2 != nil {
		return errors.New("failed in updating followee")
	}

	return nil
}
