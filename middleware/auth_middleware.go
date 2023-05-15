package middleware

import (
	"golang-api-pzm/helper"
	"golang-api-pzm/model/web"
	"net/http"
)

//membutuhkan stuct dengan kontrak handler
//karena middleware harus berbentuk handler

type AuthMiddleware struct {
	//untuk meneruskan ke handler selanjutnya
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

//func (middleware *AuthMiddleware) ServerHTTP(writer http.ResponseWriter, request *http.Request) {

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//1.mengeceknya dari header
	//2. spesifikasi nama headernya X-API-KEY
	//3. kalau misal api keynya "RAHASIA" maka ok
	//4. Kalau bukan artinya error(Unauthorize)
	if "RAHASIA" == request.Header.Get("X-API-KEY") {
		//ok
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		//error
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}
		helper.WriteToResponseBody(writer, webResponse)
	}
}
