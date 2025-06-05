package main

import (
	"log"
	"net/http"
	"os"

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
	log.SetOutput(os.Stdout)
	godotenv.Load()
	db := app.NewDB()
	validate := validator.New()

	// User
	userRepository := repository.NewUserRepository()
	userService := services.NewUserService(userRepository, db, validate)
	userController := controllers.NewUserController(userService)

	// Post
	postRepository := repository.NewPostRepository(db)
	postService := services.NewPostService(postRepository, db, validate)
	postController := controllers.NewPostController(postService)

	router := httprouter.New()

	router.POST("/auth/register", userController.Register)
	router.POST("/auth/login", userController.Login)
	router.POST("/posts", middleware.ProtectedRoute(postController.Create))
	router.GET("/posts", middleware.ProtectedRoute(postController.FindAll))
	router.GET("/users/:username", middleware.ProtectedRoute(userController.FindByUsername))

	router.PanicHandler = exception.ErrorHandler

	handler := helper.CORSMiddleware(router)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: handler,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
