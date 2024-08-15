package usecase

import (
	"blog_api/domain"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostUsecase interface {
	CreatePost(domain.Post) error
	FindPostById(string) (domain.Post, error)
	FindAllPost() ([]domain.Post, error)
	UpdatePost(domain.Post) error
	DeletePost(string) error
	LikePost(string, string) error
	UnlikePost(string, string) error
}

type postUsecase struct {
	postRepo domain.PostRepository
}

func NewPostUsecase(pr domain.PostRepository) PostUsecase {
	return &postUsecase{
		postRepo: pr,
	}
}

func (pu *postUsecase) CreatePost(post domain.Post) error {

	
	if post.Content == "" {
		return errors.New("Invalid Content")
	}
	if post.Author.Username == "" {
		return errors.New("Invalid Author ID")
	}

	objectID := primitive.NewObjectID()
	post.ID = objectID

	err := pu.postRepo.CreatePost(post)
	if err != nil {
		return err
	}

	return nil
}

func (pu *postUsecase) FindPostById(id string) (domain.Post, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	post, er := pu.postRepo.FindPostById(objID)
	if er != nil {
		return domain.Post{}, er
	}
	return post, nil
}

func (pu *postUsecase) FindAllPost() ([]domain.Post, error) {
	posts, err := pu.postRepo.FindAllPosts()
	if err != nil {
		return []domain.Post{}, err
	}

	return posts, nil
}

func (pu *postUsecase) UpdatePost(post domain.Post) error {
	err := pu.postRepo.UpdatePost(post)
	if err != nil {
		return err
	}
	return nil
}

func (pu *postUsecase) DeletePost(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	err = pu.postRepo.DeletePost(objID)
	if err != nil {
		return err
	}

	return nil
}

func (pu *postUsecase) LikePost(postID string, username string) error {
	objID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		log.Fatal(err)
	}
	var user domain.User
	user.Username = username
	err = pu.postRepo.LikePost(objID, user)
	return nil
}

func (pu *postUsecase) UnlikePost(postID string, username string) error {
	objID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		log.Fatal(err)
	}
	var user domain.User
	user.Username = username
	err = pu.postRepo.UnlikePost(objID, user)
	return nil
}
