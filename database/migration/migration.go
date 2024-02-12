package migration

import (
	"log"
	"online-store-application/database"
	"online-store-application/model/entity"
)

func InitMigration() {
	err := database.DB.AutoMigrate(&entity.User{}, &entity.Category{}, &entity.Product{})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database migrated")
}
