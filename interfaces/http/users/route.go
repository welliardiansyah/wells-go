package users

import (
	"wells-go/application/usecases"
	"wells-go/infrastructure/config"
	"wells-go/infrastructure/middleware"
	"wells-go/infrastructure/persistence"
	"wells-go/util/security"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteUsers(db *gorm.DB, router *gin.RouterGroup, cfg *config.Config, maker security.Maker) {
	repo := persistence.NewUserRepository(db)
	repoRole := persistence.NewRoleRepositoryImpl(db)
	repoAccess := persistence.NewRouteAccessRepositoryImpl(db)
	usecase := usecases.NewUserUsecase(repo, repoRole, cfg, maker)
	controller := NewUserController(usecase)

	// Public routes
	router.POST("/api/v1/users/register", controller.Register)
	router.POST("/api/v1/users/login", controller.Login)

	// Protected routes (JWT middleware)
	protected := router.Group("/api/v1/users")
	protected.Use(middleware.AuthMiddleware(maker))
	protected.Use(middleware.RoleAndPermissionMiddlewareDynamic(repoAccess))

	protected.GET("/", controller.GetUsers)
	protected.GET("/:id", controller.GetUserByID)
	protected.PUT("/:id", controller.UpdateUser)
	protected.DELETE("/:id", controller.DeleteUser)
	protected.POST("/logout", controller.Logout)
}
