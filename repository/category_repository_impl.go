package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-api-pzm/helper"
	"golang-api-pzm/model/domain"

	_ "github.com/lib/pq"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "insert into categories(id, name) values($1,$2);"
	_, err := tx.ExecContext(ctx, SQL, category.Id, category.Name)
	//dari helper
	helper.PanicIfError(err)

	// id, err := result.LastInsertId()
	// if err != nil {
	// 	panic(err)
	// }

	// category.Id = int(id)
	return category
}
func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "Update categories set name = $2 where id = $1;"
	_, err := tx.ExecContext(ctx, SQL, category.Id, category.Name)
	helper.PanicIfError(err)

	return category
}
func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "Delete from categories where id = $1;"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfError(err)
}
func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := "select id, name from categories where id = $1;"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category is not found")
	}

}
func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "select id,name from categories;"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}
