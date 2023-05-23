package admin

import (
	"net/http"
	"strconv"

	"github.com/bagusyanuar/go_deltomed_document/common"
	usecase "github.com/bagusyanuar/go_deltomed_document/usecase/admin"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService usecase.UserService
}

func NewUserController(userService usecase.UserService) UserController {
	return UserController{UserService: userService}
}

func (controller UserController) RegisterRoute(route *gin.Engine) {
	api := route.Group("/api/admin")
	{
		api.GET("/user", controller.FindAll)
	}
}

func (controller UserController) FindAll(c *gin.Context) {
	q := c.Query("q")
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	data, err := controller.UserService.FindAll(q, limit, offset)
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