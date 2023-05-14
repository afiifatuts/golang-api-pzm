package service

import (
	"context"
	"golang-api-pzm/model/web"
)

// ketika request tidak boleh memakai domain karena itu ranahnya repository
// maka kita membuat model baru representasi dari requestnya
// return nya tidak disarankan dengan domain karena nnti ada data sensitif yang bisa diekspos
// maka dari itu membuat model baru untuk response
type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}
