package controller

import (
	"blog_api/domain"
	"blog_api/usecase"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostController interface {
	CreatePost(c *gin.Context)
	UpdatePost(c *gin.Context)
	DeletePost(c *gin.Context)
	FindPostById(c *gin.Context)
	FindAllPosts(c *gin.Context)
	LikePost(c *gin.Context)
	UnlikePost(c *gin.Context)
}

type postController struct {
	postUsecase usecase.PostUsecase
}

func NewPostController(pu usecase.PostUsecase) PostController {
	return &postController{
		postUsecase: pu,
	}
}

func (pc *postController) CreatePost(c *gin.Context) {
	var post domain.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Set the Author manually using data from the context
	username := c.GetString("username")
	userID := c.GetString("user_id")
	fmt.Println("username: ", username)
	fmt.Println("userID: ", userID)
	primitiveID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	user := domain.User{
		ID:       primitiveID,
		Username: username,
		Role:     c.GetString("role"),
		Password: "",
	}
	post.Author = user

	// Set the CreatedAt and UpdatedAt times
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	// Call the use case to create the post
	err = pc.postUsecase.CreatePost(post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post created successfully"})
}

func (pc *postController) UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var post domain.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	err := pc.postUsecase.UpdatePost(id, post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully"})

}

func (pc *postController) DeletePost(c *gin.Context) {
	id := c.Param("id")
	err := pc.postUsecase.DeletePost(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}

func (pc *postController) FindPostById(c *gin.Context) {
	id := c.Param("id")
	post, err := pc.postUsecase.FindPostById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}

func (pc *postController) FindAllPosts(c *gin.Context) {
	posts, err := pc.postUsecase.FindAllPost()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)
}

func (pc *postController) LikePost(c *gin.Context) {
	postID := c.Param("post_id")
	username := c.GetString("user_id")
	err := pc.postUsecase.LikePost(postID, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post liked successfully"})
}

func (pc *postController) UnlikePost(c *gin.Context) {
	postID := c.Param("post_id")
	username := c.GetString("user_id")
	err := pc.postUsecase.UnlikePost(postID, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post unliked successfully"})
}
