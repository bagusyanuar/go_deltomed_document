package builder

import (
	adminController "github.com/bagusyanuar/go_deltomed_document/http/controller/admin"
	adminRepository "github.com/bagusyanuar/go_deltomed_document/repositories/admin"
	adminService "github.com/bagusyanuar/go_deltomed_document/service/admin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BuildGroup(route *gin.Engine, db *gorm.DB) {
	adminUserRepository := adminRepository.NewUserRepository(db)
	adminUserService := adminService.NewUserService(adminUserRepository)
	adminController := adminController.NewUserController(adminUserService)
	adminController.RegisterRoute(route)
}