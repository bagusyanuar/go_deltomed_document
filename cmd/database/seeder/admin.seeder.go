package seeder

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/bagusyanuar/go_deltomed_document/cmd/database/scheme"
	"github.com/bagusyanuar/go_deltomed_document/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func AdminSeeder(database *gorm.DB) {
	if database.Migrator().HasTable(&scheme.User{}) {
		if err := database.Where("JSON_SEARCH(roles, 'all', 'administrator') IS NOT NULL").First(&model.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			hash, err := bcrypt.GenerateFromPassword([]byte("administrator"), 13)
			if err != nil {
				panic("failed to create seeder")
			}
			password := string(hash)

			role, _ := json.Marshal([]string{"administrator"})
			admin := model.User{
				Email:    "administrator@gmail.com",
				Username: "administrator",
				Password: &password,
				Roles:    role,
			}
			if err := database.Create(&admin).Error; err != nil {
				panic("failed to create seeder")
			}
			log.Println("success create admin seeder")
		} else {
			log.Println("admin not seeded")
		}
	} else {
		log.Println("tabel users doesn't exists")
	}
}
