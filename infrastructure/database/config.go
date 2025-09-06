package database

import (
	"fmt"
	"wells-go/domain/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var gormDB *gorm.DB

func ConnectGorm(user, password, host, dbname string, port int) error {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
		host, user, password, dbname, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	gormDB = db
	err = gormDB.AutoMigrate(&entities.UserEntity{})
	if err != nil {
		return err
	}

	return nil
}

func GetGormDB() *gorm.DB {
	return gormDB
}
