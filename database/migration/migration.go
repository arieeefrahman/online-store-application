package migration

import (
	"log"
	"online-store-application/database"
	"online-store-application/model/entity"
)

func InitMigration() {
	err := database.DB.AutoMigrate(&entity.User{})

	if err != nil {
		log.Println(err)
	}

	log.Println("Database migrated")
}
