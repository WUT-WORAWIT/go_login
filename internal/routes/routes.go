package routes

import (
	"go_login/internal/handlers"
	"go_login/internal/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Public routes
	public := r.Group("/api")
	{
		public.GET("/", handlers.GetRemark)
		public.GET("/branches", handlers.GetBranch)
		public.POST("/users", handlers.CreateUser)
		public.POST("/login", handlers.Login)
	}

	// Protected routes
	protected := r.Group("/api")
	protected.Use(middleware.JWTMiddleware())
	{
		protected.GET("/users", handlers.GetUsersAll)
		protected.GET("/users/:id", handlers.GetUserByID)
		protected.PUT("/users/:id", handlers.UpdateUser)
		protected.DELETE("/users/:id", handlers.DeleteUser)
	}

	return r
}
