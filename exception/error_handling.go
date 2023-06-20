package exception

import (
	"EchoEdyP/RESTfulAPI-Clean-Architecture/helper"
	"EchoEdyP/RESTfulAPI-Clean-Architecture/models/request_response"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ErrorHandling(err error, c echo.Context) {
	if notFoundError(err, c) {
		return
	}

	if validationError(err, c) {
		return
	}

	internalServerError(err, c)
}

func notFoundError(err error, c echo.Context) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		c.Response().Header().Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
		c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")

		apiResponse := request_response.Response{
			Status:  http.StatusNotFound,
			Message: "not found",
			Data:    exception.Error(),
		}
		c.Response().WriteHeader(apiResponse.Status)
		helper.WriteToResponseBody(c, apiResponse)
		return true
	} else {
		return false
	}
}

func validationError(err error, c echo.Context) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		c.Response().Header().Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
		c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")

		apiResponse := request_response.Response{
			Status:  http.StatusBadRequest,
			Message: "bad request",
			Data:    exception.Error(),
		}
		c.Response().WriteHeader(apiResponse.Status)
		helper.WriteToResponseBody(c, apiResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(err error, c echo.Context) {
	c.Response().Header().Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")

	apiResponse := request_response.Response{
		Status:  http.StatusInternalServerError,
		Message: "internal server error",
		Data:    err.Error(),
	}
	c.Response().WriteHeader(apiResponse.Status)
	helper.WriteToResponseBody(c, apiResponse)
}
