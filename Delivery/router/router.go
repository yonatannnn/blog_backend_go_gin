package router

import (
	"blog_api/delivery/controller"
	"blog_api/infrastructure"

	"github.com/gin-gonic/gin"
)

func SetupRouter(pc controller.PostController, uc controller.UserController, md infrastructure.Middleware) *gin.Engine {
	r := gin.Default()

	// Public routes
	r.POST("auth/register", func(c *gin.Context) {
		_ = uc.Register(c)
	})
	r.POST("auth/login", func(c *gin.Context) {
		_, _ = uc.Login(c)
	})

	// Protected routes (Requires JWT Authentication)
	auth := r.Group("/")
	auth.Use(md.JWTMiddleware()) // Add your JWT middleware here
	{
		auth.GET("/users/:username", func(c *gin.Context) {
			_, _ = uc.FindUserByUsername(c)
		})

		auth.PUT("/users/:id", func(c *gin.Context) {
			_ = uc.UpdateUser(c)
		})
		auth.DELETE("/users/:username", func(c *gin.Context) {
			_ = uc.DeleteUser(c)
		})
		auth.POST("/users/follow/:id", func(c *gin.Context) {
			_ = uc.FollowUser(c)
		})

		auth.GET("/posts/:id", pc.FindPostById)
		auth.GET("/posts", pc.FindAllPosts)

		auth.POST("/posts/:post_id/like", pc.LikePost)
		auth.POST("/posts/:post_id/unlike", pc.UnlikePost)

		// Admin only routes (Requires Admin Role)
		auth.Use(md.AdminOnly()) // Add your Admin-only middleware here
		{
			auth.GET("/users", func(c *gin.Context) {
				_, _ = uc.FindAllUser(c)
			})
			auth.POST("/posts", pc.CreatePost)
			auth.PUT("/posts/:id", pc.UpdatePost)
			auth.DELETE("/posts/:id", pc.DeletePost)
		}
	}

	return r
}
