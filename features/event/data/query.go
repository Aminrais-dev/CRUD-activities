package data

import (
	"project/e-commerce/features/event"

	"gorm.io/gorm"
)

type activitiesData struct {
	DB *gorm.DB
}

func New(conn *gorm.DB) event.DataInterface {
	return &activitiesData{
		DB: conn,
	}
}

func (repo *activitiesData) SelectAll() ([]event.Core, error) {

	var data []Activities
	tx := repo.DB.Find(&data)
	if tx.Error != nil {
		return nil, tx.Error
	}

	dataAll := toCoreList(data)

	return dataAll, nil
}

func (repo *activitiesData) SelectById(param int) (event.Core, error) {

	var data Activities
	tx := repo.DB.First(&data, "id = ?", param)
	if tx.Error != nil {
		return event.Core{}, tx.Error
	}

	return data.toCore(), nil
}

func (repo *activitiesData) CreateData(data event.Core) (event.Core, error) {
	dataModel := fromCore(data)
	tx := repo.DB.Create(&dataModel)
	if tx.Error != nil {
		return event.Core{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return event.Core{}, tx.Error
	}

	var dataReturn Activities
	repo.DB.Where("activity_type = ? AND institution = ? AND when = ? AND remarks = ? AND objective = ? ", dataModel.ActivityType, dataModel.Institution, dataModel.When, dataModel.Remarks, dataModel.Objective).First(&dataReturn)

	return dataReturn.toCore(), nil

}

func (repo *activitiesData) UpdateData(newData event.Core) (event.Core, error) {

	dataModel := fromCore(newData)

	tx := repo.DB.Model(&Activities{}).Where("id = ? ", newData.ID).Updates(dataModel)
	if tx.Error != nil {
		return event.Core{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return event.Core{}, tx.Error
	}

	var data Activities
	repo.DB.First(&data, "id = ? ", newData.ID)

	return data.toCore(), nil
}
