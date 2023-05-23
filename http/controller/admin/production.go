package admin

import (
	"net/http"
	"strconv"

	"github.com/bagusyanuar/go_deltomed_document/common"
	usecase "github.com/bagusyanuar/go_deltomed_document/usecase/admin"
	"github.com/gin-gonic/gin"
)

type ProductionController struct {
	ProductionService usecase.ProductionService
}

func NewProductionController(productionService usecase.ProductionService) ProductionController {
	return ProductionController{ProductionService: productionService}
}

func (controller *ProductionController) RegisterRoute(route *gin.Engine) {
	api := route.Group("/api/admin")
	{
		api.GET("/production", controller.FindAll)
	}
}

func (controller *ProductionController) FindAll(c *gin.Context) {
	q := c.Query("q")
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	data, err := controller.ProductionService.FindAll(q, limit, offset)
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
