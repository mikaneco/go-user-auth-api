package database

import (
	"goapi/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(setting string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(setting), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Println("Database connected")
	return db, nil
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		return err
	}
	log.Println("Database migrated")
	return nil
}
