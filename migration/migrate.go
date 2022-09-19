package migration

import (
	event "project/e-commerce/features/event/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&event.Activities{})

}
