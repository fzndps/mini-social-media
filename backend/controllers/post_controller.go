package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PostController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
