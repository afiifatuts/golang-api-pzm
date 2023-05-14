package controller

import (
	"golang-api-pzm/helper"
	"golang-api-pzm/model/web"
	"golang-api-pzm/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryContollerImpl struct {
	CategoryService service.CategoryService
}

//parameter category service
//return controller
//yang diekpose controllerinterface
//yang direturn controllerimplementation

// kegunaan kita membuat interface dulu
// kalau di OOP mirip Polymorphishm
func NewCategoryController(categoryService service.CategoryService) CategoryContoller {
	return &CategoryContollerImpl{
		CategoryService: categoryService,
	}
}

// bisa Encode dan Decode jadi tidak usah conversi ke string
func (controller *CategoryContollerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//mengambil request dari body
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryContollerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &categoryUpdateRequest)

	//mengambil params
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryUpdateRequest.Id = id

	categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *CategoryContollerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	//mengambil params
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.CategoryService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *CategoryContollerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//mengambil params
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryContollerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//mengambil params
	categoryResponses := controller.CategoryService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
