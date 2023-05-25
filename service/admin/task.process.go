package admin

import (
	"github.com/bagusyanuar/go_deltomed_document/http/request"
	"github.com/bagusyanuar/go_deltomed_document/model"
	usecase "github.com/bagusyanuar/go_deltomed_document/usecase/admin"
)

type implementsTaskProcessService struct {
	TaskProcessRepository usecase.TaskProcessRepository
}

// Create implements admin.TaskProcessService
func (service *implementsTaskProcessService) Create(request request.CreateTaskProcessRequest) (data *model.TaskProcess, err error) {
	// taskID, err := uuid.Parse(request.TaskID)
	// if err != nil {
	// 	return nil, err
	// }

	// productionStepID, err := uuid.Parse(request.ProductionStepID)
	// if err != nil {
	// 	return nil, err
	// }

	// productionSubStepID, err := uuid.Parse(request.ProductionSubStepID)
	// if err != nil {
	// 	return nil, err
	// }

	// image := new(string)
	// if request.Image != nil {
	// 	if _, err := os.Stat(common.ImagePath); os.IsNotExist(err) {
	// 		err = os.MkdirAll(common.ImagePath, os.ModePerm)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 	}
	// 	ext := filepath.Ext(request.Image.Filename)
	// 	fileName := fmt.Sprintf("%s/%s%s", common.ImagePath, uuid.New().String(), ext)
	// 	image = &fileName
	// 	common.UploadFile(request.Image, fileName)
	// }
	// now := time.Now()
	// entity := model.TaskProcess{
	// 	TaskID:              taskID,
	// 	ProductionStepID:    productionStepID,
	// 	ProductionsubStepID: productionSubStepID,
	// 	SubmittedAt:         &now,
	// 	Image:               image,
	// }
	// return service.TaskProcessRepository.Create(entity)
	panic("unimplemented")
}

// Delete implements admin.TaskProcessService
func (service *implementsTaskProcessService) Delete(id string) (err error) {
	panic("unimplemented")
}

// FindAll implements admin.TaskProcessService
func (service *implementsTaskProcessService) FindAll(param string, limit int, offset int) (data []model.TaskProcess, err error) {
	panic("unimplemented")
}

// FindByID implements admin.TaskProcessService
func (service *implementsTaskProcessService) FindByID(id string) (data *model.TaskProcess, err error) {
	panic("unimplemented")
}

// Patch implements admin.TaskProcessService
func (service *implementsTaskProcessService) Patch(id string, request request.CreateTaskProcessRequest) (data *model.TaskProcess, err error) {
	panic("unimplemented")
}

func NewTaskProcessService(taskProcessRepository usecase.TaskProcessRepository) usecase.TaskProcessService {
	return &implementsTaskProcessService{
		TaskProcessRepository: taskProcessRepository,
	}
}
