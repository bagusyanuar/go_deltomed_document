package seeder

import "gorm.io/gorm"

func Seed(database *gorm.DB) {
	AdminSeeder(database)
	SupportSeeder(database)
}
