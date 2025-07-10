package usecases

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"go_auth/databases/models"
	"go_auth/databases/repositories"
	"go_auth/handlers/http/payloads/request"
	"go_auth/lib/util"
)

type (
	UserUseCase interface {
		GetListUsers(query request.ListUserRequest) (users []models.Users, totalRow int64, err error)
		RegisterUser(data models.Users) (user models.Users, err error)
		LoginUser(request request.LoginRequest) (token string, err error)
	}

	UserUseCaseImpl struct {
		userRepo repositories.UserRepository
	}
)

func NewUserUsecase(
	userRepo repositories.UserRepository,
) UserUseCase {
	return &UserUseCaseImpl{
		userRepo: userRepo,
	}
}

func (u *UserUseCaseImpl) GetListUsers(query request.ListUserRequest) (users []models.Users, totalRow int64, err error) {
	users, totalRow, err = u.userRepo.ListUsers(query)
	if err != nil {
		return
	}
	return
}

func (u *UserUseCaseImpl) RegisterUser(data models.Users) (user models.Users, err error) {

	salt := "inisaltyagesya"

	h := sha1.New()

	if _, err := h.Write([]byte(salt + data.Password + salt)); err != nil {
		return models.Users{}, err
	}

	bs := h.Sum(nil)
	data.Password = fmt.Sprintf("%x", bs)

	user, err = u.userRepo.RegisterUser(data)
	if err != nil {
		return
	}
	return
}

func (u *UserUseCaseImpl) LoginUser(request request.LoginRequest) (token string, err error) {

	result, err := u.userRepo.FindUserByUsername(request.Username)
	if err != nil {
		return "", err
	}

	salt := "inisaltyagesya"
	h := sha1.New()
	if _, err := h.Write([]byte(salt + request.Password + salt)); err != nil {
		return "", err
	}

	bs := h.Sum(nil)
	Password := fmt.Sprintf("%x", bs)
	if result.Password != Password {
		return "", errors.New("invalid username or password")
	}

	tokenString, err := util.CreateToken(result.Username)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
