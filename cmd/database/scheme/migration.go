package scheme

import (
	"log"

	"gorm.io/gorm"
)

func Migrate(database *gorm.DB) {
	database.AutoMigrate(&User{})
	database.AutoMigrate(&Production{})
	log.Println("success migrate database")
}
