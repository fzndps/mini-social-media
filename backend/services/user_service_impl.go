package services

import (
	"context"
	"database/sql"
	"errors"

	"github.com/fzndps/mini-social-media/backend/exception"
	"github.com/fzndps/mini-social-media/backend/helper"
	"github.com/fzndps/mini-social-media/backend/models/domain"
	"github.com/fzndps/mini-social-media/backend/models/web"
	"github.com/fzndps/mini-social-media/backend/repository"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userService repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userService,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) Register(ctx context.Context, request web.UserRegisterRequest) web.UserRegisterResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	if service.UserRepository.IsUsernameExists(ctx, tx, request.Username) {
		panic(exception.NewConflictRequestError("Username already taken"))
	}

	if service.UserRepository.IsEmailExists(ctx, tx, request.Email) {
		panic(exception.NewConflictRequestError("Email already taken"))
	}

	defer helper.CommitOrRollback(tx)

	user := domain.User{
		Username: request.Username,
		Email:    request.Email,
		Password: string(hashedPassword),
	}

	user = service.UserRepository.Create(ctx, tx, user)
	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Login(ctx context.Context, request web.UserLoginRequest) (web.UserLoginResponse, error) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.LoginByUsername(ctx, tx, request.Username)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return web.UserLoginResponse{}, errors.New("invalid credentials")
	}

	token, err := helper.GenerateJWT(user.Id, user.Username)
	helper.PanicIfError(err)

	return helper.ToUserLoginResponse(user, token), nil

}

func (service *UserServiceImpl) FindByUsername(ctx context.Context, username string) web.UserRegisterResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByUsername(ctx, tx, username)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToUserResponse(user)
}
