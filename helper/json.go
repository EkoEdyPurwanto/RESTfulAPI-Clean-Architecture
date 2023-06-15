package helper

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func ReadFromRequestBody(c echo.Context, result interface{}) error {
	err := c.Bind(result)
	if err != nil {
		return err
	}
	return nil
}

func WriteToResponseBody(c echo.Context, response interface{}) error {
	return c.JSON(http.StatusOK, response)
}
