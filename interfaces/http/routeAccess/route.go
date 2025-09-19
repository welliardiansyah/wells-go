package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"wells-go/application/usecases"
	"wells-go/infrastructure/middleware"
	"wells-go/infrastructure/persistence"
	"wells-go/util/security"
)

func RouteAccessRoutes(db *gorm.DB, router *gin.RouterGroup, maker security.Maker) {
	repo := persistence.NewRouteAccessRepositoryImpl(db)
	usecase := usecases.NewRouteAccessUsecase(repo)
	controller := NewRouteAccessHandler(usecase)

	protected := router.Group("/route-access")
	protected.Use(middleware.AuthMiddleware(maker))
	protected.Use(middleware.RoleAndPermissionMiddlewareDynamic(repo))

	protected.GET("/get/all", controller.GetAll)
	protected.GET("/get/:id", controller.GetByID)
	protected.POST("/create", controller.Create)
	protected.PUT("/update/:id", controller.Update)
	protected.DELETE("/delete/:id", controller.Delete)
	protected.GET("/get/by-role", controller.GetAllByRole)
	protected.POST("/get/by-name", controller.GetAllByName)
}
