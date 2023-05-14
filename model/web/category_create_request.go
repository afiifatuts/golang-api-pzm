package web

type CategoryCreateRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required"`
}
