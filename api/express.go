package api

import (
	"epaygo/core/helper"
	"tp-api-go/service"

	"net/http"
	. "tp-api-go/core"
	. "tp-api-go/core/dto"

	"github.com/labstack/echo"
)

func CreateExpressOrder(c echo.Context) error {
	expressInfoDto := new(ExpressInfoDto)
	if err := c.Bind(expressInfoDto); err != nil {
		return c.JSON(http.StatusBadRequest, helper.NewApiMessage(10004, err.Error(), "Object"))
	} else {

		if _, apiError := service.NewExpressOrderService().CreateExpressOrder(expressInfoDto); apiError != nil {
			return c.JSON(http.StatusInternalServerError, APIResult{Success: false, Error: apiError})
		} else {
			return c.JSON(http.StatusCreated, APIResult{Success: true, Result: nil})
		}
	}

}
