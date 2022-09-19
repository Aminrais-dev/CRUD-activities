package usecase

import (
	"project/e-commerce/features/event"
)

type activitiesUsecase struct {
	activitiesData event.DataInterface
}

func New(data event.DataInterface) event.UsecaseInterface {
	return &activitiesUsecase{
		activitiesData: data,
	}
}

func (usecase *activitiesUsecase) GetAll() ([]event.Core, error) {
	data, err := usecase.activitiesData.SelectAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (usecase *activitiesUsecase) GetById(id int) (event.Core, error) {
	data, err := usecase.activitiesData.SelectById(id)
	if err != nil {
		return event.Core{}, err
	}
	return data, nil
}

func (usecase *activitiesUsecase) PostData(data event.Core) int {
	if data.ActivityType == "" || data.Institution == "" || data.Objective == "" || data.Remarks == "" || data.When == "" {
		return -1
	}

	row := usecase.activitiesData.CreateData(data)
	if row == -1 {
		return -1
	}

	return row
}

func (usecase *activitiesUsecase) PutData(data event.Core) int {
	row := usecase.activitiesData.UpdateData(data)
	if row == -1 {
		return -1
	}

	return row

}
