package migrate

import (
	"fmt"
	"wells-go/domain/entities"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	models := []interface{}{
		&entities.UserEntity{},
		&entities.RoleEntity{},
		&entities.PermissionEntity{},
		&entities.RouteAccessEntities{},
		&entities.PathRouteEntities{},
	}

	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			return err
		}
		fmt.Printf("✅ AutoMigrate berhasil untuk model: %T\n", model)
	}

	fmt.Println("🚀 Semua migrasi berhasil!")
	return nil
}
