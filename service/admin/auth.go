package admin

import (
	"github.com/bagusyanuar/go_deltomed_document/common"
	"github.com/bagusyanuar/go_deltomed_document/http/request"
	"github.com/bagusyanuar/go_deltomed_document/model"
	usecase "github.com/bagusyanuar/go_deltomed_document/usecase/admin"
	"golang.org/x/crypto/bcrypt"
)

type implementsAuthService struct {
	AuthRepository usecase.AuthRepository
}

// SignIn implements admin.AuthService
func (service *implementsAuthService) SignIn(request request.CreateSignInRequest) (accessToken string, err error) {
	entity := model.User{
		Username: request.Username,
	}
	user, err := service.AuthRepository.SignIn(entity)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(request.Password))
	if err != nil {
		return "", common.ErrorPasswordNotMatch
	}

	jwtSign := common.JWTSignReturn{
		ID:       user.ID,
		Username: user.Username,
		Role:     "administrator",
	}
	return common.GenerateAccessToken(&jwtSign)
}

func NewAuthService(authRepository usecase.AuthRepository) usecase.AuthService {
	return &implementsAuthService{
		AuthRepository: authRepository,
	}
}
