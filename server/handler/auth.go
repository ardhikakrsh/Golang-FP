package handler

import (
	"github.com/biskitsx/Leave-Management-System/service"
	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	service service.AuthService
}

func NewAuthHandler(service service.AuthService) authHandler {
	return authHandler{
		service: service,
	}
}

func (h *authHandler) Signup(c *fiber.Ctx) error {
	req := service.NewSignupRequest{}
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	res, err := h.service.Signup(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.JSON(res)
}

func (h *authHandler) Login(c *fiber.Ctx) error {
	req := service.LoginRequest{}
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	res, err := h.service.Login(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	accessCookie := &fiber.Cookie{
		Name:   "access_token",
		Value:  res.Token,
		MaxAge: 600,
	}
	c.Cookie(accessCookie)
	return c.JSON(res)
}

func (h *authHandler) Logout(c *fiber.Ctx) error {
	accessCookie := &fiber.Cookie{
		Name:   "access_token",
		Value:  "",
		MaxAge: -1,
	}
	c.Cookie(accessCookie)
	return c.SendString("Logged out successfully")
}
