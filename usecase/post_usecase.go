package usecase

import (
	"blog_api/domain"
	"errors"
)

type PostUsecase interface {
	CreatePost(domain.Post) error
	FindPostById(string) (domain.Post, error)
	FindAllPost() ([]domain.Post, error)
	UpdatePost(domain.Post) error
	DeletePost(string) error
	LikePost(string, string) error
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
	if post.ID == "" {
		return errors.New("Invalid Post ID")
	}
	if post.Content == "" {
		return errors.New("Invalid Content")
	}
	if post.Author.Username == "" {
		return errors.New("Invalid Author ID")
	}
	return pu.postRepo.CreatePost(post)
}

func (pu *postUsecase) FindPostById(id string) (domain.Post, error) {
	
	return domain.Post{}, nil
}

func (pu *postUsecase) FindAllPost() ([]domain.Post, error) {
	
	return []domain.Post{}, nil
}

func (pu *postUsecase) UpdatePost(post domain.Post) error {
	
	return nil
}

func (pu *postUsecase) DeletePost(id string) error {
	
	return nil
}

func (pu *postUsecase) LikePost(postID string, userID string) error {
	// TODO: Implement LikePost method
	return nil
}


