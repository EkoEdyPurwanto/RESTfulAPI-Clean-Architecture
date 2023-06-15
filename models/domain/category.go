package domain

import (
	"EchoEdyP/RESTfulAPI-Clean-Architecture/models/request_response"
	"context"
	"database/sql"
)

type Category struct {
	Id   int
	Name string
}

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category Category) (Category, error)
	Update(ctx context.Context, tx *sql.Tx, category Category) (Category, error)
	Delete(ctx context.Context, tx *sql.Tx, category Category) error
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]Category, error)
}

type CategoryUseCase interface {
	Create(ctx context.Context, request request_response.CategoryCreateRequest) (response request_response.CategoryResponse, err error)
	Update(ctx context.Context, request request_response.CategoryUpdateRequest) (response request_response.CategoryResponse, err error)
	Delete(ctx context.Context, categoryId int) (err error)
	FindById(ctx context.Context, categoryId int) (response request_response.CategoryResponse, err error)
	FindAll(ctx context.Context) (response []request_response.CategoryResponse, err error)
}
