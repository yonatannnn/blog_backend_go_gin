package controller

import (
	"blog_api/domain"
	"blog_api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Register(*gin.Context) error
	Login(*gin.Context) (domain.User, error)
	FindUserByUsername(*gin.Context) (domain.User, error)
	FindAllUser(*gin.Context) ([]domain.User, error)
	UpdateUser(*gin.Context) error
	DeleteUser(*gin.Context) error
	FollowUser(*gin.Context) error
	FindUserByID(*gin.Context) (domain.User, error)
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

func (uc *userController) FindUserByUsername(c *gin.Context) (domain.User, error) {
	username := c.Param("username")
	user, err := uc.userUsecase.FindUserByUsername(username)
	if err != nil {
		c.JSON(500, err.Error())
		return domain.User{}, err
	}
	c.JSON(http.StatusOK, user)
	return user, nil
}

func (uc *userController) FindUserByID(c *gin.Context) (domain.User, error) {
	id := c.Param("id")
	user, err := uc.userUsecase.FindUserById(id)
	if err != nil {
		c.JSON(500, err.Error())
		return domain.User{}, err
	}
	c.JSON(http.StatusOK, user)
	return user, nil
}

func (uc *userController) FindAllUser(c *gin.Context) ([]domain.User, error) {
	users, err := uc.userUsecase.FindAllUser()
	if err != nil {
		c.JSON(500, err.Error())
		return []domain.User{}, err
	}
	c.JSON(http.StatusOK, users)
	return users, nil
}

func (uc *userController) UpdateUser(c *gin.Context) error {
	id := c.Param("id")
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return nil
	}
	

	uc.userUsecase.UpdateUser(id, user)
	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
	})
	return nil
}

func (uc *userController) DeleteUser(c *gin.Context) error {
	username := c.Param("username")
	err := uc.userUsecase.DeleteUser(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
	return nil
}

func (uc *userController) FollowUser(c *gin.Context) error {
	foloweeID := c.Param("id")
	folowerID := c.GetString("user_id")
	err := uc.userUsecase.FollowUser(folowerID, foloweeID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	return nil
}
