package usecase

import (
	"EchoEdyP/RESTfulAPI-Clean-Architecture/helper"
	"EchoEdyP/RESTfulAPI-Clean-Architecture/models/domain"
	"EchoEdyP/RESTfulAPI-Clean-Architecture/models/request_response"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type CategoryUseCase struct {
	CategoryRepository domain.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func (useCase *CategoryUseCase) Create(ctx context.Context, request request_response.CategoryCreateRequest) (response request_response.CategoryResponse, err error) {
	err = useCase.Validate.Struct(request)
	if err != nil {
		return request_response.CategoryResponse{}, err
	}

	tx, err := useCase.DB.Begin()
	if err != nil {
		return request_response.CategoryResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category, err = useCase.CategoryRepository.Save(ctx, tx, category)
	if err != nil {
		return request_response.CategoryResponse{}, err
	}

	return helper.ToCategoryRespones(category), nil
}

func (useCase *CategoryUseCase) Update(ctx context.Context, request request_response.CategoryUpdateRequest) (response request_response.CategoryResponse, err error) {
	err = useCase.Validate.Struct(request)
	if err != nil {
		return request_response.CategoryResponse{}, err
	}

	tx, err := useCase.DB.Begin()
	if err != nil {
		return request_response.CategoryResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	category, err := useCase.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		return request_response.CategoryResponse{}, err
	}

	category.Name = request.Name

	category, err = useCase.CategoryRepository.Update(ctx, tx, category)
	if err != nil {
		return request_response.CategoryResponse{}, err
	}

	return helper.ToCategoryRespones(category), nil
}

func (useCase *CategoryUseCase) Delete(ctx context.Context, categoryId int) (err error) {
	tx, err := useCase.DB.Begin()
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	category, err := useCase.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		return err
	}

	err = useCase.CategoryRepository.Delete(ctx, tx, category)
	if err != nil {
		return err
	}

	return nil
}

func (useCase *CategoryUseCase) FindById(ctx context.Context, categoryId int) (response request_response.CategoryResponse, err error) {
	tx, err := useCase.DB.Begin()
	if err != nil {
		return request_response.CategoryResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	category, err := useCase.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		return request_response.CategoryResponse{}, err
	}

	return helper.ToCategoryRespones(category), nil
}

func (useCase *CategoryUseCase) FindAll(ctx context.Context) (response []request_response.CategoryResponse, err error) {
	tx, err := useCase.DB.Begin()
	if err != nil {
		return []request_response.CategoryResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	categories, err := useCase.CategoryRepository.FindAll(ctx, tx)
	if err != nil {
		return []request_response.CategoryResponse{}, err
	}

	return helper.ToCategoryResponses(categories), nil
}
