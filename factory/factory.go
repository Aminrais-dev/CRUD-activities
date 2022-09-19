package factory

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	eventData "project/e-commerce/features/event/data"
	eventDelivery "project/e-commerce/features/event/delivery"
	eventUsecase "project/e-commerce/features/event/usecase"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	eventDataFactory := eventData.New(db)
	eventUsecaseFactory := eventUsecase.New(eventDataFactory)
	eventDelivery.New(e, eventUsecaseFactory)

}
