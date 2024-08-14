package controller

import (
	"blog_api/domain"
	"blog_api/usecase"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Register(*gin.Context) error
	Login(*gin.Context) (domain.User, error)
	FindUserById(string) (domain.User, error)
	FindAllUser() ([]domain.User, error)
	UpdateUser(domain.User) error
	DeleteUser(string) error
	FollowUser(string, string) error
}

type userController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController(uu usecase.UserUsecase) UserController {
	return &userController{
		userUsecase: uu,
	}
}

func (uc *userController) Register(c *gin.Context) error {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, err.Error())
		return err
	}

	err := uc.userUsecase.CreateUser(user)
	if err != nil {
		c.JSON(500, err.Error())
		return err
	}

	c.JSON(200, gin.H{"message": "User created successfully"})
	return nil
}

func (uc *userController) Login(c *gin.Context) (domain.User, error) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, err.Error())
		return domain.User{}, err
	}

	user, err, token := uc.userUsecase.Login(user.Username, user.Password)
	if err != nil {
		c.JSON(500, err.Error())
		return domain.User{}, err
	}

	c.JSON(200, gin.H{"message": "User logged in successfully", "token": token})
	return user, nil
}

func (uc *userController) FindUserById(id string) (domain.User, error) {

	return domain.User{}, nil
}

func (uc *userController) FindAllUser() ([]domain.User, error) {
	// TODO: Implement FindAllUser method
	return []domain.User{}, nil
}

func (uc *userController) UpdateUser(user domain.User) error {
	// TODO: Implement UpdateUser method
	return nil
}

func (uc *userController) DeleteUser(id string) error {
	// TODO: Implement DeleteUser method
	return nil
}

func (uc *userController) FollowUser(followerID string, followeeID string) error {
	// TODO: Implement FollowUser method
	return nil
}
