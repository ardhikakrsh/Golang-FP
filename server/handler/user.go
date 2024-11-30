package handler

import (
	"github.com/biskitsx/Leave-Management-System/service"
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	GetUsers(c *fiber.Ctx) error
	GetMe(c *fiber.Ctx) error
}

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{
		userService: userService,
	}
}

func (h *userHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.userService.GetUsers()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.JSON(users)
}

func (h *userHandler) GetMe(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)
	return c.JSON(fiber.Map{
		"userId": userId,
	})
}
