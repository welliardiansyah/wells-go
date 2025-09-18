package pathRoute

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"wells-go/application/usecases"
	"wells-go/infrastructure/config"
	"wells-go/infrastructure/middleware"
	"wells-go/infrastructure/persistence"
	"wells-go/util/security"
)

func RoutePathRoute(db *gorm.DB, router *gin.RouterGroup, cfg *config.Config, maker security.Maker) {
	repo := persistence.NewPathRouteRepositoryImpl(db)
	repoAccess := persistence.NewRouteAccessRepositoryImpl(db)
	usecase := usecases.NewPathRouteUsecase(repo)
	controller := NewPathRouteHandler(usecase)

	// Protected route group
	protected := router.Group("/api/v1/path-route")
	protected.Use(middleware.AuthMiddleware(maker))
	protected.Use(middleware.RoleAndPermissionMiddlewareDynamic(repoAccess))

	protected.GET("", controller.GetAllRoutes)
}
