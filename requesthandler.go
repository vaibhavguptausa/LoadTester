package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func registerHandlers(e *echo.Echo) {
	e.POST("/load", loadTest)
}

func loadTest(c echo.Context) error {
	var loadReq = new(LoadRequest)
	if err := c.Bind(loadReq); err != nil {
		return c.JSON(http.StatusBadRequest, JsonResponse{Success: false, Message: err.Error()})
	}

	if err := c.Validate(loadReq); err != nil {
		return c.JSON(http.StatusBadRequest, JsonResponse{Success: false, Message: err.Error()})
	}

	go InitLoad(*loadReq)
	return c.JSON(http.StatusAccepted, "Request accepted")
}
