package admin

import (
	"encoding/json"

	"github.com/bagusyanuar/go_deltomed_document/http/request"
	"github.com/bagusyanuar/go_deltomed_document/model"
	usecase "github.com/bagusyanuar/go_deltomed_document/usecase/admin"
	"golang.org/x/crypto/bcrypt"
)

type implementsUserService struct {
	UserRepository usecase.UserRepository
}

// Create implements admin.UserService
func (service *implementsUserService) Create(request request.CreateUserRequest) (data *model.User, err error) {
	roles, _ := json.Marshal([]string{request.Roles})
	var password string
	if request.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), 13)
		if err != nil {
			return nil, err
		}
		password = string(hash)
	}
	entity := model.User{
		Email:      request.Email,
		Username:   request.Username,
		Password:   &password,
		Roles:      roles,
	}
	return service.UserRepository.Create(entity)
}

// Delete implements admin.UserService
func (service *implementsUserService) Delete(id string) (err error) {
	return service.UserRepository.Delete(id)
}

// FindAll implements admin.UserService
func (service *implementsUserService) FindAll(param string, limit int, offset int) (data []model.User, err error) {
	return service.UserRepository.FindAll(param, limit, offset)
}

// FindByID implements admin.UserService
func (service *implementsUserService) FindByID(id string) (data *model.User, err error) {
	return service.UserRepository.FindByID(id)
}

// Patch implements admin.UserService
func (service *implementsUserService) Patch(id string, request request.CreateUserRequest) (data *model.User, err error) {
	entity := model.User{
		Email:    request.Email,
		Username: request.Username,
	}
	return service.UserRepository.Patch(id, entity)
}

func NewUserService(userRepository usecase.UserRepository) usecase.UserService {
	return &implementsUserService{
		UserRepository: userRepository,
	}
}
