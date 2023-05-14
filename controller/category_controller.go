package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CategoryContoller interface {
	//parameter mengikuti dari http handler
	//karena memakai httprouter harus pakai params
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
