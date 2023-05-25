package support

import (
	"net/http"
	"strconv"

	"github.com/bagusyanuar/go_deltomed_document/common"
	"github.com/bagusyanuar/go_deltomed_document/http/middleware"
	usecase "github.com/bagusyanuar/go_deltomed_document/usecase/support"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
	TaskService usecase.TaskService
	Middleware  middleware.Middleware
}

func NewTaskController(taskService usecase.TaskService) TaskController {
	return TaskController{TaskService: taskService, Middleware: middleware.Middleware{}}
}

func (controller *TaskController) RegisterRoute(route *gin.Engine) {
	api := route.Group("/api/support")
	{
		api.GET("/task", controller.Middleware.Auth(), controller.GetData)
	}
}

func catch(c *gin.Context) {
	if r := recover(); r != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: "internal server error",
			Data:    nil,
		})
		return
	}
}
func (controller *TaskController) GetData(c *gin.Context) {
	defer catch(c)
	authorizedUser := c.MustGet("user").(*common.JWTClaims)
	authorizedID := authorizedUser.Unique.String()
	q := c.Query("q")
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	data, err := controller.TaskService.GetData(authorizedID, q, limit, offset)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, common.APIResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}
