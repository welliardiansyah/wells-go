package role

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"wells-go/application/usecases"
	"wells-go/infrastructure/config"
	"wells-go/infrastructure/middleware"
	"wells-go/infrastructure/persistence"
	"wells-go/util/security"
)

func RouteRoles(db *gorm.DB, router *gin.RouterGroup, cfg *config.Config, maker security.Maker) {
	repo := persistence.NewRoleRepositoryImpl(db)
	repoPermission := persistence.NewPermissionRepositoryImpl(db)
	repoAccess := persistence.NewRouteAccessRepositoryImpl(db)

	usecase := usecases.NewRoleUsecase(repo, repoPermission)
	controller := NewRoleController(usecase)

	protected := router.Group("/roles")
	protected.Use(middleware.AuthMiddleware(maker))
	protected.Use(middleware.RoleAndPermissionMiddlewareDynamic(repoAccess))

	protected.POST("/create/role", controller.CreateRole)
	protected.GET("/get/all/role", controller.GetAllRoles)
	protected.PUT("/update/:id", controller.UpdateRole)
	protected.DELETE("/delete/:id", controller.DeleteRole)
	protected.GET("/get/:id", controller.GetRoleByID)
}
