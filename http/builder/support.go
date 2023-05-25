package builder

import (
	controller "github.com/bagusyanuar/go_deltomed_document/http/controller/support"
	repository "github.com/bagusyanuar/go_deltomed_document/repositories/support"
	service "github.com/bagusyanuar/go_deltomed_document/service/support"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SupportBuilder(route *gin.Engine, db *gorm.DB) {
	authRepository := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepository)
	authController := controller.NewAuthController(authService)
	authController.RegisterRoute(route)

	taskRepository := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepository)
	taskController := controller.NewTaskController(taskService)
	taskController.RegisterRoute(route)
}
