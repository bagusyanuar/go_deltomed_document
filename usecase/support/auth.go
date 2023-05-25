package support

import (
	"github.com/bagusyanuar/go_deltomed_document/http/request"
	"github.com/bagusyanuar/go_deltomed_document/model"
)

type AuthRepository interface {
	SignIn(user model.User) (data *model.User, err error)
}

type AuthService interface {
	SignIn(request request.CreateSignInRequest) (accessToken string, err error)
}
