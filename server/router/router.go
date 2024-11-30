package router

import (
	"github.com/biskitsx/Leave-Management-System/handler"
	"github.com/biskitsx/Leave-Management-System/middleware"
	"github.com/biskitsx/Leave-Management-System/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Init(app *fiber.App, db *gorm.DB) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	setUserRoute(app, db)
	setAuthRoute(app, db)
	setLeaveRoute(app, db)
}

func setUserRoute(app *fiber.App, db *gorm.DB) {
	userHandler := handler.NewUserHandler(service.NewUserService(db))
	app.Get("/users", userHandler.GetUsers)
	app.Get("/users/me", middleware.VerifyLogin, userHandler.GetMe)
}

func setAuthRoute(app *fiber.App, db *gorm.DB) {
	authHandler := handler.NewAuthHandler(service.NewAuthService(db))
	app.Post("/auth/signup", authHandler.Signup)
	app.Post("/auth/login", authHandler.Login)
	app.Post("/auth/logout", authHandler.Logout)
}

func setLeaveRoute(app *fiber.App, db *gorm.DB) {
	leaveHandler := handler.NewLeaveHandler(service.NewLeaveService(db))
	app.Get("/leaves", leaveHandler.GetLeaves)
	app.Post("/leaves", middleware.VerifyLogin, leaveHandler.AddLeave)
	app.Get("/leaves/me", middleware.VerifyLogin, leaveHandler.GetUserLeaves)
	// app.Put("/leaves/:id", leaveHandler.UpdateLeave)
	// app.Delete("/leaves/:id", leaveHandler.DeleteLeave)
}
