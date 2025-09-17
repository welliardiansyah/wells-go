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
	usecase := usecases.NewRoleUsecase(repo, repoPermission)
	controller := NewRoleController(usecase)

	protected := router.Group("/roles")
	protected.Use(middleware.AuthMiddleware(maker))

	protected.POST("/create/role", middleware.RoleAndPermissionMiddleware([]string{"Admin"}, []string{"create"}), controller.CreateRole)

	protected.GET("/get/all/role",
		middleware.RoleAndPermissionMiddleware(
			[]string{"Admin", "Manager"},
			[]string{"read"},
		),
		controller.GetAllRoles,
	)

	protected.PUT("/update/:id",
		middleware.RoleAndPermissionMiddleware(
			[]string{"Admin"},
			[]string{"update"},
		),
		controller.UpdateRole,
	)

	protected.DELETE("/delete/:id",
		middleware.RoleAndPermissionMiddleware(
			[]string{"Admin"},
			[]string{"delete"},
		),
		controller.DeleteRole,
	)

	protected.GET("/get/:id",
		middleware.RoleAndPermissionMiddleware(
			[]string{"Admin", "Manager", "User"},
			[]string{"view"},
		),
		controller.GetRoleByID,
	)
}
