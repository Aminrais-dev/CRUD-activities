package delivery

import (
	"project/e-commerce/features/event"
	"project/e-commerce/utils/helper"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type ActivitiesDelivery struct {
	activitiesUsecase event.UsecaseInterface
}

func New(e *echo.Echo, usecase event.UsecaseInterface) {
	handler := &ActivitiesDelivery{
		activitiesUsecase: usecase,
	}

	e.GET("/activities", handler.GetAllData)
	e.GET("/activities/:id", handler.GetByIdData)
	e.POST("/activities", handler.PostData)
	e.PUT("activities/:id", handler.PutDataAct)

}

func (delivery *ActivitiesDelivery) PostData(c echo.Context) error {

	var dataProduct Request
	errBind := c.Bind(&dataProduct)
	if errBind != nil {
		return c.JSON(400, "error bind")
	}

	data, err := delivery.activitiesUsecase.PostData(dataProduct.toCoreAct())
	if err != nil {
		return c.JSON(400, "error insert data")
	}
	return c.JSON(201, data)

}

func (delivery *ActivitiesDelivery) GetAllData(c echo.Context) error {

	data, err := delivery.activitiesUsecase.GetAll()
	if err != nil {
		return c.JSON(400, "failed get all data")
	}

	return c.JSON(200, map[string]interface{}{
		"activities": data,
	})

}

func (delivery *ActivitiesDelivery) GetByIdData(c echo.Context) error {

	id := c.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(400, "param must be number")
	}

	data, errGet := delivery.activitiesUsecase.GetById(ID)
	if errGet != nil {
		return c.JSON(400, helper.FailedResponseHelper("error get by id data"))
	}

	return c.JSON(200, data)

}

func (delivery *ActivitiesDelivery) PutDataAct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("param must be number"))
	}

	var dataUpdate Request
	errBind := c.Bind(&dataUpdate)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("error bind data"))
	}

	var add event.Core
	if dataUpdate.ActivityType != "" {
		add.ActivityType = dataUpdate.ActivityType
	}
	if dataUpdate.Institution != "" {
		add.Institution = dataUpdate.Institution
	}
	if dataUpdate.Objective != "" {
		add.Objective = dataUpdate.Objective
	}
	if dataUpdate.Remarks != "" {
		add.Remarks = dataUpdate.Remarks
	}
	if dataUpdate.When == "" {
		date, _ := time.Parse(time.RFC3339, dataUpdate.When)
		add.When = date
	}

	add.ID = uint(id)

	data, errUpd := delivery.activitiesUsecase.PutData(add)
	if errUpd != nil {
		return c.JSON(400, helper.FailedResponseHelper("error update data"))
	}

	return c.JSON(200, data)
}
