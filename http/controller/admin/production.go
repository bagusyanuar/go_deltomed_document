package admin

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/bagusyanuar/go_deltomed_document/common"
	"github.com/bagusyanuar/go_deltomed_document/http/request"
	usecase "github.com/bagusyanuar/go_deltomed_document/usecase/admin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
		api.POST("/production", controller.Create)
		api.GET("/production/:id", controller.FindByID)
		api.PATCH("/production/:id", controller.Patch)
		api.DELETE("/production/:id", controller.Delete)
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

func (controller *ProductionController) FindByID(c *gin.Context) {
	id := c.Param("id")
	data, err := controller.ProductionService.FindByID(id)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			c.JSON(http.StatusOK, common.APIResponse{
				Code:    http.StatusOK,
				Message: err.Error(),
				Data:    data,
			})
			return
		}
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

func (controller *ProductionController) Create(c *gin.Context) {
	var request request.CreateProductionRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, common.APIResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	data, err := controller.ProductionService.Create(request)
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

func (controller *ProductionController) Patch(c *gin.Context) {
	id := c.Param("id")
	var request request.CreateProductionRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, common.APIResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	data, err := controller.ProductionService.Patch(id, request)
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

func (controller *ProductionController) Delete(c *gin.Context) {
	id := c.Param("id")
	err := controller.ProductionService.Delete(id)
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
