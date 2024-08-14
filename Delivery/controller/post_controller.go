package controller

import (
	"blog_api/domain"
	"blog_api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostController interface {
	CreatePost(c *gin.Context)
	UpdatePost(c *gin.Context)
	DeletePost(c *gin.Context)
	FindPostById(c *gin.Context)
	FindAllPosts(c *gin.Context)
	LikePost(c *gin.Context)
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := pc.postUsecase.CreatePost(post)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post created successfully"})
}

func (pc *postController) UpdatePost(c *gin.Context) {
	// TODO: Implement UpdatePost method
}

func (pc *postController) DeletePost(c *gin.Context) {
	// TODO: Implement DeletePost method
}

func (pc *postController) FindPostById(c *gin.Context) {
	// TODO: Implement FindPostById method
}

func (pc *postController) FindAllPosts(c *gin.Context) {
	// TODO: Implement FindAllPosts method
}

func (pc *postController) LikePost(c *gin.Context) {
	// TODO: Implement LikePost method
}
