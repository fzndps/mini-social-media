package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/fzndps/mini-social-media/backend/app"
	"github.com/fzndps/mini-social-media/backend/controllers"
	commentCtrl "github.com/fzndps/mini-social-media/backend/controllers/comment"
	"github.com/fzndps/mini-social-media/backend/exception"
	"github.com/fzndps/mini-social-media/backend/helper"
	"github.com/fzndps/mini-social-media/backend/middleware"
	"github.com/fzndps/mini-social-media/backend/repository"
	commentRepo "github.com/fzndps/mini-social-media/backend/repository/comment"
	"github.com/fzndps/mini-social-media/backend/services"
	commentSrvc "github.com/fzndps/mini-social-media/backend/services/comment"
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

	// User Profile
	userPostRepository := repository.NewUserPostRepository(db)
	userPostService := services.NewUserPostService(userPostRepository, db)
	userPostController := controllers.NewUserPostController(userPostService)

	// Post
	postRepository := repository.NewPostRepository(db)
	postService := services.NewPostService(postRepository, db, validate)
	postController := controllers.NewPostController(postService)

	// Comment
	commentRepository := commentRepo.NewCommentRepository(db)
	commentService := commentSrvc.NewCommentService(commentRepository, postRepository, userRepository, db, validate)
	commentController := commentCtrl.NewCommentController(commentService)

	router := httprouter.New()

	// users
	router.POST("/auth/register", userController.Register)
	router.POST("/auth/login", userController.Login)
	router.GET("/api/users/profile/:userId", middleware.ProtectedRoute(userPostController.UserPostProfile))
	router.POST("/api/posts/:postId/comments", middleware.ProtectedRoute(commentController.CreateComment)) // postId bisa diambil dari body atau path
	router.GET("/api/posts/:postId/comments", middleware.ProtectedRoute(commentController.FindPostWithCommentsById))
	router.PUT("/api/comments/:commentId", middleware.ProtectedRoute(commentController.UpdateComment))
	router.DELETE("/api/posts/:postId/comments/:commentId", middleware.ProtectedRoute(commentController.Delete))

	// router.GET("/users/:username", middleware.ProtectedRoute(userController.FindByUsername))

	// post
	router.POST("/api/posts", middleware.ProtectedRoute(postController.Create))
	router.GET("/api/posts", middleware.ProtectedRoute(postController.FindAll))

	router.PanicHandler = exception.ErrorHandler

	handler := helper.CORSMiddleware(router)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: handler,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
