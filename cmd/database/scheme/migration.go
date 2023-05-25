package scheme

import (
	"log"

	"gorm.io/gorm"
)

func Migrate(database *gorm.DB) {
	database.AutoMigrate(&User{})
	database.AutoMigrate(&Production{})
	database.AutoMigrate(&ProductionStep{})
	database.AutoMigrate(&ProductionSubStep{})
	database.AutoMigrate(&Task{})
	database.AutoMigrate(&TaskProcess{})
	log.Println("success migrate database")
}
