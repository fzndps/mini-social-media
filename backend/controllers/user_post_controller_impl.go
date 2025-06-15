package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/fzndps/mini-social-media/backend/helper"
	"github.com/fzndps/mini-social-media/backend/models/web"
	"github.com/fzndps/mini-social-media/backend/services"
	"github.com/julienschmidt/httprouter"
)

type UserPostControllerImpl struct {
	UserPostService services.UserPostService
}

func NewUserPostController(userPostService services.UserPostService) UserPostController {
	return &UserPostControllerImpl{
		UserPostService: userPostService,
	}
}

func (controller *UserPostControllerImpl) UserPostProfile(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userPostResponse := controller.UserPostService.UserPostProfile(request.Context(), id)
	log.Println("Isi userPostResponse :", userPostResponse)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userPostResponse,
	}
	log.Println("Isi webResponse :", webResponse)
	helper.WriteResponseBody(writer, webResponse)
}
