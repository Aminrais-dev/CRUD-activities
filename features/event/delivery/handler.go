package delivery

import (
	"project/e-commerce/features/event"
	"project/e-commerce/utils/helper"
	"strconv"

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
		return c.JSON(400, helper.FailedResponseHelper("error bind"))
	}

	row := delivery.activitiesUsecase.PostData(dataProduct.toCoreAct())
	if row == -1 {
		return c.JSON(400, helper.FailedResponseHelper("error insert data"))
	}
	return c.JSON(201, helper.SuccessResponseHelper("success insert data"))

}

func (delivery *ActivitiesDelivery) GetAllData(c echo.Context) error {

	data, err := delivery.activitiesUsecase.GetAll()
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("error get all data"))
	}

	return c.JSON(200, helper.SuccessDataResponseHelper("success get all data", toRespon(data)))

}

func (delivery *ActivitiesDelivery) GetByIdData(c echo.Context) error {

	id := c.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("param must be number"))
	}

	data, errGet := delivery.activitiesUsecase.GetById(ID)
	if errGet != nil {
		return c.JSON(400, helper.FailedResponseHelper("error get by id data"))
	}

	return c.JSON(200, helper.SuccessDataResponseHelper("success get all data", data))

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
		add.When = dataUpdate.When
	}

	add.ID = uint(id)

	row := delivery.activitiesUsecase.PutData(add)
	if row == -1 {
		return c.JSON(400, helper.FailedResponseHelper("error update data"))
	}

	return c.JSON(200, helper.SuccessResponseHelper("Success update data"))
}
