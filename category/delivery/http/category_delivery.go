package http

import (
	"EchoEdyP/RESTfulAPI-Clean-Architecture/exception"
	"EchoEdyP/RESTfulAPI-Clean-Architecture/helper"
	"EchoEdyP/RESTfulAPI-Clean-Architecture/models/domain"
	"EchoEdyP/RESTfulAPI-Clean-Architecture/models/request_response"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type CategoryDelivery struct {
	CategoryUseCase domain.CategoryUseCase
}

func NewCategoryDelivery(e *echo.Echo, categoryUseCase domain.CategoryUseCase) {
	delivery := &CategoryDelivery{
		CategoryUseCase: categoryUseCase,
	}
	e.GET("/category", delivery.FindAll)
	e.POST("/category", delivery.Create)
	e.PUT("/category/:categoryId", delivery.Update)
	e.GET("/category/:categoryId", delivery.FindById)
	e.DELETE("/category/:categoryId", delivery.Delete)

	e.HTTPErrorHandler = exception.ErrorHandling
}

func (delivery *CategoryDelivery) Create(c echo.Context) error {
	categoryCreateRequest := request_response.CategoryCreateRequest{}
	err := helper.ReadFromRequestBody(c, &categoryCreateRequest)
	if err != nil {
		return err
	}

	categoryResponse, err := delivery.CategoryUseCase.Create(c.Request().Context(), categoryCreateRequest)
	if err != nil {
		return err
	}
	response := request_response.Response{
		Status:  http.StatusOK,
		Message: "OK",
		Data:    categoryResponse,
	}

	return helper.WriteToResponseBody(c, response)
}

func (delivery *CategoryDelivery) Update(c echo.Context) error {
	categoryUpdateRequest := request_response.CategoryUpdateRequest{}
	err := helper.ReadFromRequestBody(c, &categoryUpdateRequest)
	if err != nil {
		return err
	}

	categoryId := c.Param("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryUpdateRequest.Id = id

	categoryResponse, err := delivery.CategoryUseCase.Update(c.Request().Context(), categoryUpdateRequest)
	if err != nil {
		return err
	}
	response := request_response.Response{
		Status:  http.StatusOK,
		Message: "OK",
		Data:    categoryResponse,
	}

	return helper.WriteToResponseBody(c, response)
}

func (delivery *CategoryDelivery) Delete(c echo.Context) error {
	categoryId := c.Param("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	delivery.CategoryUseCase.Delete(c.Request().Context(), id)
	response := request_response.Response{
		Status:  http.StatusOK,
		Message: "OK",
	}

	return helper.WriteToResponseBody(c, response)
}

func (delivery *CategoryDelivery) FindById(c echo.Context) error {
	categoryId := c.Param("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryResponse, err := delivery.CategoryUseCase.FindById(c.Request().Context(), id)
	if err != nil {
		return err
	}
	response := request_response.Response{
		Status:  http.StatusOK,
		Message: "OK",
		Data:    categoryResponse,
	}

	return helper.WriteToResponseBody(c, response)
}

func (delivery *CategoryDelivery) FindAll(c echo.Context) error {
	categoryResponses, err := delivery.CategoryUseCase.FindAll(c.Request().Context())
	if err != nil {
		return err
	}
	response := request_response.Response{
		Status:  http.StatusOK,
		Message: "OK",
		Data:    categoryResponses,
	}

	return helper.WriteToResponseBody(c, response)
}
