package service

import (
	"context"
	"database/sql"
	"golang-api-pzm/exception"
	"golang-api-pzm/helper"
	"golang-api-pzm/model/domain"
	"golang-api-pzm/model/web"
	"golang-api-pzm/repository"

	"github.com/go-playground/validator/v10"
)

// kita butuh repository
// karena nnti manipulasi datanya menggunakan repository
type CategoryServiceImpl struct {
	//interface
	CategoryRepository repository.CategoryRepository
	//menggunakan pointer karena struct bukan interface
	DB       *sql.DB
	Validate *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	//1. memulai transaction dimulai dari service karena kalau di toko online biasanya ditambah ada inventory dll
	//2. kemudian mengirim ke repository
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	//validation commit
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Id:   request.Id,
		Name: request.Name,
	}

	//menyimpan data
	category = service.CategoryRepository.Save(ctx, tx, category)

	//conversi dari model category ke categoryresponse
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	// memulai transaction
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	// validation commit
	defer helper.CommitOrRollback(tx)

	//divalidasi apakah id nya ada atau tidak

	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	//menggunakan exeption func not found
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	//mengubah nama category
	category.Name = request.Name

	//menyimpan data
	category = service.CategoryRepository.Update(ctx, tx, category)

	//conversi dari model category ke categoryresponse
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	// validation commit
	defer helper.CommitOrRollback(tx)

	//divalidasi apakah id nya ada atau tidak

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	//menggunakan exeption func not found
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	// memulai transaction
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	// validation commit
	defer helper.CommitOrRollback(tx)

	//divalidasi apakah id nya ada atau tidak

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	//menggunakan exeption func not found
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	// memulai transaction
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	// validation commit
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)
}
