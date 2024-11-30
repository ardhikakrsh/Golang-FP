package handler

import (
	"github.com/biskitsx/Leave-Management-System/service"
	"github.com/gofiber/fiber/v2"
)

type leaveHandler struct {
	leaveService service.LeaveService
}

func NewLeaveHandler(leaveService service.LeaveService) *leaveHandler {
	return &leaveHandler{
		leaveService: leaveService,
	}
}

func (h *leaveHandler) AddLeave(c *fiber.Ctx) error {
	req := service.AddLeaveRequest{}
	userId := c.Locals("userId").(uint)
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	res, err := h.leaveService.AddLeave(req, userId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.JSON(res)
}

func (h *leaveHandler) GetUserLeaves(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)
	leaves, err := h.leaveService.GetLeavesByUser(userId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.JSON(leaves)
}

func (h *leaveHandler) GetLeaves(c *fiber.Ctx) error {
	leaves, err := h.leaveService.GetLeaves()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.JSON(leaves)
}
