package repository

import (
	"EchoEdyP/RESTfulAPI-Clean-Architecture/models/domain"
	"context"
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
)

type CategoryRepository struct {
}

func NewCategoryRepository() domain.CategoryRepository {
	return &CategoryRepository{}
}

func (repository *CategoryRepository) Save(ctx context.Context, tx *sql.Tx, category domain.Category) (domain.Category, error) {
	SQL := "INSERT INTO category(name) VALUES ($1) RETURNING id"
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	if err != nil {
		return domain.Category{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.Category{}, err
	}

	category.Id = int(id)
	return category, nil

}

func (repository *CategoryRepository) Update(ctx context.Context, tx *sql.Tx, category domain.Category) (domain.Category, error) {
	SQL := "UPDATE category SET name=$1 WHERE id=$2"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	if err != nil {
		return domain.Category{}, err
	}

	return category, nil

}

func (repository *CategoryRepository) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) error {
	SQL := "DELETE FROM category WHERE id=$1"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	if err != nil {
		return err
	}

	return nil
}

func (repository *CategoryRepository) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := "SELECT id, name FROM category WHERE id=$1"
	row := tx.QueryRowContext(ctx, SQL, categoryId)

	var category domain.Category
	err := row.Scan(&category.Id, &category.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Category{}, errors.New("Category is not found")
		}
		return domain.Category{}, err
	}

	return category, nil
}

func (repository *CategoryRepository) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Category, error) {
	SQL := "SELECT id, name FROM category"
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}
