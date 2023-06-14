package helper

import (
	"EchoEdyP/RESTfulAPI-CleanArch/models/domain"
	"EchoEdyP/RESTfulAPI-CleanArch/models/request_response"
)

func ToCategoryRespones(category domain.Category) request_response.CategoryResponse {
	return request_response.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []request_response.CategoryResponse {
	var categoryResponses []request_response.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryRespones(category))
	}
	return categoryResponses
}
