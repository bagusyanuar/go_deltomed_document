package builder

import (
	controller "github.com/bagusyanuar/go_deltomed_document/http/controller/admin"
	repository "github.com/bagusyanuar/go_deltomed_document/repositories/admin"
	service "github.com/bagusyanuar/go_deltomed_document/service/admin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AdminBuilder(route *gin.Engine, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	userController.RegisterRoute(route)

	productionRepository := repository.NewProductionRepository(db)
	productionService := service.NewProductionService(productionRepository)
	productionController := controller.NewProductionController(productionService)
	productionController.RegisterRoute(route)

	productionStepRepository := repository.NewProductionStepRepository(db)
	productionStepService := service.NewProductionStepService(productionStepRepository)
	productionStepController := controller.NewProductionStepController(productionStepService)
	productionStepController.RegisterRoute(route)

	productionSubStepRepository := repository.NewProductionSubStepRepository(db)
	productionSubStepService := service.NewProductionSubStepService(productionSubStepRepository)
	productionSubStepController := controller.NewProductionSubStepController(productionSubStepService)
	productionSubStepController.RegisterRoute(route)

	taskRepository := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepository)
	taskController := controller.NewTaskController(taskService)
	taskController.RegisterRoute(route)
}
