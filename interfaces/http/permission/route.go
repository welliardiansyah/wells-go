package permission

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"wells-go/application/usecases"
	"wells-go/infrastructure/config"
	"wells-go/infrastructure/middleware"
	"wells-go/infrastructure/persistence"
	"wells-go/util/security"
)

func RoutePermissions(db *gorm.DB, router *gin.RouterGroup, cfg *config.Config, maker security.Maker) {
	repo := persistence.NewPermissionRepositoryImpl(db)
	repoAccess := persistence.NewRouteAccessRepositoryImpl(db)
	usecase := usecases.NewPermissionUsecase(repo)
	controller := NewPermissionController(usecase)

	protected := router.Group("/api/v1/permissions")
	protected.Use(middleware.AuthMiddleware(maker))
	protected.Use(middleware.RoleAndPermissionMiddlewareDynamic(repoAccess))

	protected.POST("/create", controller.Create)
	protected.GET("/get/all", controller.FindAll)
	protected.GET("/get/:id", controller.FindByID)
	protected.PUT("/update/:id", controller.Update)
	protected.DELETE("/delete/:id", controller.Delete)
}
