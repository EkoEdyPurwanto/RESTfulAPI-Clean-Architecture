package repository

import (
	"EchoEdyP/RESTfulAPI-CleanArch/models"
	"context"
	"database/sql"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category models.Category) (models.Category, error)
	Update(ctx context.Context, tx *sql.Tx, category models.Category) (models.Category, error)
	Delete(ctx context.Context, tx *sql.Tx, category models.Category) error
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (models.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]models.Category, error)
}
