package usecase

import (
	"EchoEdyP/RESTfulAPI-CleanArch/models/request_response"
	"context"
)

type CategoryUseCase interface {
	Create(ctx context.Context, request request_response.CategoryCreateRequest) (response request_response.CategoryResponse, err error)
	Update(ctx context.Context, request request_response.CategoryUpdateRequest) (response request_response.CategoryResponse, err error)
	Delete(ctx context.Context, categoryId int) (err error)
	FindById(ctx context.Context, categoryId int) (response request_response.CategoryResponse, err error)
	FindAll(ctx context.Context) (response []request_response.CategoryResponse, err error)
}
