package admin

import (
	"net/http"
	"strconv"

	"github.com/bagusyanuar/go_deltomed_document/common"
	"github.com/bagusyanuar/go_deltomed_document/http/request"
	usecase "github.com/bagusyanuar/go_deltomed_document/usecase/admin"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
	TaskService usecase.TaskService
}

func NewTaskController(taskService usecase.TaskService) TaskController {
	return TaskController{TaskService: taskService}
}

func (controller *TaskController) RegisterRoute(route *gin.Engine) {
	api := route.Group("/api/admin")
	{
		api.GET("/task", controller.FindAll)
		api.POST("/task", controller.Create)
		api.GET("/task/:id", controller.FindByID)
		api.PATCH("/task/:id", controller.Patch)
		api.DELETE("/task/:id", controller.Delete)
	}
}

func (controller *TaskController) FindAll(c *gin.Context) {
	q := c.Query("q")
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	data, err := controller.TaskService.FindAll(q, limit, offset)
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

func (controller *TaskController) FindByID(c *gin.Context) {
	id := c.Param("id")
	data, err := controller.TaskService.FindByID(id)
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

func (controller *TaskController) Create(c *gin.Context) {
	var request request.CreateTaskRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, common.APIResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	data, err := controller.TaskService.Create(request)
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

func (controller *TaskController) Patch(c *gin.Context) {
	id := c.Param("id")
	var request request.CreateTaskRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, common.APIResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	data, err := controller.TaskService.Patch(id, request)
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

func (controller *TaskController) Delete(c *gin.Context) {
	id := c.Param("id")
	err := controller.TaskService.Delete(id)
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
		Data:    nil,
	})
}
