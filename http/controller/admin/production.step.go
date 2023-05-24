package admin

import (
	"net/http"
	"strconv"

	"github.com/bagusyanuar/go_deltomed_document/common"
	"github.com/bagusyanuar/go_deltomed_document/http/request"
	usecase "github.com/bagusyanuar/go_deltomed_document/usecase/admin"
	"github.com/gin-gonic/gin"
)

type ProductionStepController struct {
	ProductionStepService usecase.ProductionStepService
}

func NewProductionStepController(productionStepService usecase.ProductionStepService) ProductionStepController {
	return ProductionStepController{ProductionStepService: productionStepService}
}

func (controller *ProductionStepController) RegisterRoute(route *gin.Engine) {
	api := route.Group("/api/admin")
	{
		api.GET("/production-step", controller.FindAll)
		api.POST("/production-step", controller.Create)
		api.GET("/production-step/:id", controller.FindByID)
		api.PATCH("/production-step/:id", controller.Patch)
		api.DELETE("/production-step/:id", controller.Delete)
	}
}

func (controller *ProductionStepController) FindAll(c *gin.Context) {
	q := c.Query("q")
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	data, err := controller.ProductionStepService.FindAll(q, limit, offset)
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

func (controller *ProductionStepController) FindByID(c *gin.Context) {
	id := c.Param("id")
	data, err := controller.ProductionStepService.FindByID(id)
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

func (controller *ProductionStepController) Create(c *gin.Context) {
	var request request.CreateProductionStepRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, common.APIResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	data, err := controller.ProductionStepService.Create(request)
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

func (controller *ProductionStepController) Patch(c *gin.Context) {
	id := c.Param("id")
	var request request.CreateProductionStepRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, common.APIResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	data, err := controller.ProductionStepService.Patch(id, request)
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

func (controller *ProductionStepController) Delete(c *gin.Context) {
	id := c.Param("id")
	err := controller.ProductionStepService.Delete(id)
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
