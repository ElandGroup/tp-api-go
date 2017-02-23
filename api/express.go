package api

import (
	"tp-api-go/service"

	"net/http"
	. "tp-api-go/core"
	. "tp-api-go/core/dto"

	"github.com/labstack/echo"
)

func CreateExpressOrder(c echo.Context) error {

	expressInfoDto := new(ExpressInfoDto)
	if resp, apiError := service.NewExpressOrderService().CreateExpressOrder(expressInfoDto); apiError != nil {
		return c.JSON(http.StatusInternalServerError, APIResult{Success: true, Error: apiError})
	} else {
		return c.JSON(http.StatusOK, APIResult{Success: true, Result: resp})
	}
}
