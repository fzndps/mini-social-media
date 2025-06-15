package comment

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CommentController interface {
	CreateComment(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindPostWithCommentsById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateComment(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
