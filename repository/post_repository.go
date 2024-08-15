package repository

import (
	"blog_api/domain"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type postRepository struct {
	databaseCollection *mongo.Collection
	context            context.Context
}

func NewPostRepository(databaseCollection *mongo.Collection) domain.PostRepository {
	return &postRepository{
		databaseCollection: databaseCollection,
		context:            context.TODO(),
	}
}

func (r *postRepository) CreatePost(post domain.Post) error {
	_, err := r.databaseCollection.InsertOne(r.context, post)
	if err != nil {
		return errors.New("Failed to create post")
	}
	return nil
}

func (r *postRepository) FindPostById(id primitive.ObjectID) (domain.Post, error) {
	var post domain.Post
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	err := r.databaseCollection.FindOne(r.context, filter).Decode(&post)
	if err != nil {
		return domain.Post{}, errors.New("Post not found")
	}

	return post, nil
}

func (r *postRepository) FindAllPosts() ([]domain.Post, error) {
	var posts []domain.Post
	cursor, err := r.databaseCollection.Find(r.context, bson.D{})
	if err != nil {
		return []domain.Post{}, errors.New("Failed to fetch posts")
	}

	for cursor.Next(r.context) {
		var post domain.Post
		err := cursor.Decode(&post)
		if err != nil {
			return []domain.Post{}, errors.New("Failed to fetch posts")
		}
		posts = append(posts, post)
	}

	return []domain.Post{}, nil
}

func (r *postRepository) UpdatePost(post domain.Post) error {
	filter := bson.D{bson.E{Key: "_id", Value: post.ID}}
	updatedFields := bson.D{}

	if post.Content != "" {
		updatedFields = append(updatedFields, bson.E{Key: "content", Value: post.Content})
	}
	if post.Author.Username != "" {
		updatedFields = append(updatedFields, bson.E{Key: "author", Value: post.Author})
	}

	if post.Title != "" {
		updatedFields = append(updatedFields, bson.E{Key: "title", Value: post.Title})
	}

	update := bson.D{
		bson.E{
			Key: "$set", Value: updatedFields},
	}

	_, err := r.databaseCollection.UpdateOne(r.context, filter, update)
	if err != nil {
		return errors.New("Failed to update post")
	}

	return nil
}

func (r *postRepository) DeletePost(id primitive.ObjectID) error {
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	err := r.databaseCollection.FindOneAndDelete(r.context, filter)
	if err.Err() != nil {
		return errors.New("Failed to delete post")
	}
	return nil
}

func (r *postRepository) LikePost(post_id primitive.ObjectID, user domain.User) error {
	filter := bson.D{bson.E{Key: "_id", Value: post_id}}
	update := bson.D{
		bson.E{
			Key: "$push", Value: bson.D{bson.E{Key: "likes", Value: user}},
		},
	}
	_, err := r.databaseCollection.UpdateOne(r.context, filter, update)
	if err != nil {
		return errors.New("Failed to like post")
	}
	return nil
}

func (r *postRepository) UnlikePost(post_id primitive.ObjectID, user domain.User) error {
	filter := bson.D{bson.E{Key: "_id", Value: post_id}}
	update := bson.D{
		bson.E{
			Key: "$pull", Value: bson.D{bson.E{Key: "likes", Value: user}},
		},
	}
	_, err := r.databaseCollection.UpdateOne(r.context, filter, update)
	if err != nil {
		return errors.New("Failed to unlike post")
	}
	return nil
}
