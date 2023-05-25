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

func SupportSeeder(database *gorm.DB) {
	if database.Migrator().HasTable(&scheme.User{}) {
		if err := database.Where("JSON_SEARCH(roles, 'all', 'support') IS NOT NULL").First(&model.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			hash, err := bcrypt.GenerateFromPassword([]byte("support"), 13)
			if err != nil {
				panic("failed to create seeder")
			}
			password := string(hash)

			role, _ := json.Marshal([]string{"support"})
			support := model.User{
				Email:    "support@gmail.com",
				Username: "support",
				Password: &password,
				Roles:    role,
			}
			if err := database.Create(&support).Error; err != nil {
				panic("failed to create seeder")
			}
			log.Println("success create support seeder")
		} else {
			log.Println("support not seeded")
		}
	} else {
		log.Println("tabel users doesn't exists")
	}
}
