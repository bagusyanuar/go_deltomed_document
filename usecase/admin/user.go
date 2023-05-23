package admin

import (
	"github.com/bagusyanuar/go_deltomed_document/http/request"
	"github.com/bagusyanuar/go_deltomed_document/model"
)

type UserRepository interface {
	FindAll(param string, limit int, offset int) (data []model.User, err error)
	FindByID(id string) (data *model.User, err error)
	Create(entity model.User) (data *model.User, err error)
	Patch(id string, entity model.User) (data *model.User, err error)
	Delete(id string) (err error)
}

type UserService interface {
	FindAll(param string, limit int, offset int) (data []model.User, err error)
	FindByID(id string) (data *model.User, err error)
	Create(request request.CreateUserRequest) (data *model.User, err error)
	Patch(id string, request request.CreateUserRequest) (data *model.User, err error)
	Delete(id string) (err error)
}