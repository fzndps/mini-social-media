package controllers

import (
	"net/http"

	"github.com/fzndps/mini-social-media/backend/exception"
	"github.com/fzndps/mini-social-media/backend/helper"
	"github.com/fzndps/mini-social-media/backend/models/web"
	"github.com/fzndps/mini-social-media/backend/services"
	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := web.UserRegisterRequest{}
	helper.ReadFromRequestBody(request, &userCreateRequest)

	userResponse := controller.UserService.Register(request.Context(), userCreateRequest)

	webResponse := web.WebResponse{
		Code:   http.StatusCreated,
		Status: "CREATED",
		Data:   userResponse,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userLoginRequest := web.UserLoginRequest{}
	helper.ReadFromRequestBody(request, &userLoginRequest)

	userLoginResponse, err := controller.UserService.Login(request.Context(), userLoginRequest)
	if err != nil {
		panic(exception.NewUnauthorizedError(err.Error()))
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userLoginResponse,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) FindByUsername(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	username := params.ByName("username")

	response := controller.UserService.FindByUsername(request.Context(), username)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	}

	helper.WriteResponseBody(writer, webResponse)
}
