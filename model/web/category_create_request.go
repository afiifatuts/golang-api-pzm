package web

type CategoryCreateRequest struct {
	Id   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}
