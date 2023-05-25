package support

import (
	"net/http"

	"github.com/bagusyanuar/go_deltomed_document/common"
	"github.com/bagusyanuar/go_deltomed_document/http/request"
	usecase "github.com/bagusyanuar/go_deltomed_document/usecase/support"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	AuthService usecase.AuthService
}

func NewAuthController(authService usecase.AuthService) AuthController {
	return AuthController{AuthService: authService}
}

func (controller *AuthController) RegisterRoute(route *gin.Engine) {
	api := route.Group("/api/support")
	{
		api.POST("/sign-in", controller.SignIn)
	}
}

func (controller *AuthController) SignIn(c *gin.Context) {
	var request request.CreateSignInRequest
	c.BindJSON(&request)
	accessToken, err := controller.AuthService.SignIn(request)
	if err != nil {
		switch err {
		case common.ErrorPasswordNotMatch:
			c.AbortWithStatusJSON(http.StatusUnauthorized, common.APIResponse{
				Code:    http.StatusUnauthorized,
				Message: err.Error(),
				Data:    nil,
			})
			return
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, common.APIResponse{
				Code:    http.StatusNotFound,
				Message: err.Error(),
				Data:    nil,
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, common.APIResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
				Data:    nil,
			})
			return
		}
	}
	c.JSON(http.StatusOK, common.APIResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    accessToken,
	})
}
