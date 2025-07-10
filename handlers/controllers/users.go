package controller

import (
	"go_auth/databases/models"
	"go_auth/handlers/http/payloads/request"
	pagination "go_auth/lib"
	"go_auth/usecases"

	"github.com/gofiber/fiber/v2"
)

type (
	UserController interface {
		ListUsers(c *fiber.Ctx) error
		Register(c *fiber.Ctx) error
		Login(c *fiber.Ctx) error
	}

	UserControllerImpl struct {
		userUseCase usecases.UserUseCase
	}
)

func NewUserController(userUseCase usecases.UserUseCase) UserController {
	return &UserControllerImpl{
		userUseCase: userUseCase,
	}
}

func (ctrl *UserControllerImpl) ListUsers(c *fiber.Ctx) error {
	query := request.ListUserRequest{}
	c.QueryParser(&query)

	users, totalRow, err := ctrl.userUseCase.GetListUsers(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":          fiber.StatusInternalServerError,
			"error_message": err.Error(),
		})
	}

	paginationData := pagination.Data(users, int(totalRow), pagination.GetOffset(query.Page, query.Size))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code": fiber.StatusOK,
		"data": paginationData,
	})
}

func (ctrl *UserControllerImpl) Register(c *fiber.Ctx) error {
	request := models.Users{}

	c.BodyParser(&request)

	user, err := ctrl.userUseCase.RegisterUser(request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":          fiber.StatusInternalServerError,
			"error_message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code": fiber.StatusOK,
		"data": user,
	})
}

func (ctrl *UserControllerImpl) Login(c *fiber.Ctx) error {
	request := request.LoginRequest{}

	c.BodyParser(&request)

	token, err := ctrl.userUseCase.LoginUser(request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":          fiber.StatusInternalServerError,
			"error_message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code": fiber.StatusOK,
		"data": token,
	})
}
