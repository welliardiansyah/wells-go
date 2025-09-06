package users

import (
	"wells-go/application/usecases"
	"wells-go/infrastructure/config"
	"wells-go/infrastructure/middleware"
	"wells-go/infrastructure/persistence"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteUsers(db *gorm.DB, router *gin.RouterGroup, cfg *config.Config) {
	repo := persistence.NewUserRepository(db)
	usecase := usecases.NewUserUsecase(repo, cfg)
	controller := NewUserController(usecase)

	// Public routes
	router.POST("/users/register", controller.Register)
	router.POST("/users/login", controller.Login)

	// Protected routes (JWT middleware)
	protected := router.Group("/users")
	protected.Use(middleware.AuthMiddleware())

	protected.GET("/", controller.GetUsers)
	protected.GET("/:id", controller.GetUserByID)
	protected.PUT("/:id", controller.UpdateUser)
	protected.DELETE("/:id", controller.DeleteUser)
	protected.POST("/logout", controller.Logout)
}
