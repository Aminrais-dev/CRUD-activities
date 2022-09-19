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

func (repo *activitiesData) CreateData(data event.Core) int {
	dataModel := fromCore(data)
	tx := repo.DB.Create(&dataModel)
	if tx.Error != nil {
		return -1
	}
	return int(tx.RowsAffected)
}

func (repo *activitiesData) UpdateData(newData event.Core) int {

	dataModel := fromCore(newData)

	tx := repo.DB.Model(&Activities{}).Where("id = ? ", newData.ID).Updates(dataModel)
	if tx.Error != nil {
		return -1
	}
	if tx.RowsAffected == 0 {
		return -1
	}

	return 1
}
