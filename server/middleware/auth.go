package middleware

import (
	"fmt"

	"github.com/biskitsx/Leave-Management-System/service"
	"github.com/gofiber/fiber/v2"
)

func VerifyLogin(c *fiber.Ctx) error {
	cookie := c.Cookies("access_token")
	fmt.Println("Access Token Cookie:", cookie) // Log the cookie

	if cookie == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	jwtService := service.NewJwtService()
	decodedToken, err := jwtService.VerifyToken(cookie)
	if err != nil {
		fmt.Println("Error verifying token:", err) // Log error if token verification fails
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	userId, username, err := jwtService.ExtractAccessToken(decodedToken)
	if err != nil {
		fmt.Println("Error extracting token:", err) // Log error if extraction fails
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}
	c.Locals("userId", userId)
	c.Locals("username", username)
	return c.Next()
}
