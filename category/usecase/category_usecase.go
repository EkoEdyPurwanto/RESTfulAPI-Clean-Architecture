package usecase

import (
	"EchoEdyP/RESTfulAPI-CleanArch/models/request_response"
	"context"
)

type CategoryUseCase interface {
	Create(ctx context.Context, request request_response.CategoryCreateRequest) request_response.CategoryResponse
	Update(ctx context.Context, request request_response.CategoryUpdateRequest) request_response.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) request_response.CategoryResponse
	FindAll(ctx context.Context) []request_response.CategoryResponse
}
