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

func (usecase *activitiesUsecase) PostData(data event.Core) (event.Core, error) {
	if data.ActivityType == "" || data.Institution == "" || data.Objective == "" || data.Remarks == "" {
		return event.Core{}, nil
	}

	data, errData := usecase.activitiesData.CreateData(data)
	if errData != nil {
		return event.Core{}, errData
	}

	return data, nil
}

func (usecase *activitiesUsecase) PutData(data event.Core) (event.Core, error) {
	data, err := usecase.activitiesData.UpdateData(data)
	if err != nil {
		return event.Core{}, err
	}

	return data, nil

}
