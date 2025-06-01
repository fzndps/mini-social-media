package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/fzndps/mini-social-media/backend/app"
	"github.com/fzndps/mini-social-media/backend/controllers"
	"github.com/fzndps/mini-social-media/backend/exception"
	"github.com/fzndps/mini-social-media/backend/helper"
	"github.com/fzndps/mini-social-media/backend/middleware"
	"github.com/fzndps/mini-social-media/backend/repository"
	"github.com/fzndps/mini-social-media/backend/services"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {
	godotenv.Load()
	db := app.NewDB()
	validate := validator.New()

	userRepository := repository.NewUserRepository()
	userService := services.NewUserService(userRepository, db, validate)
	userController := controllers.NewUserController(userService)

	router := httprouter.New()

	router.POST("/auth/register", userController.Register)
	router.POST("/auth/login", userController.Login)
	router.GET("/users/:username", middleware.ProtectedRoute(userController.FindByUsername))

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
