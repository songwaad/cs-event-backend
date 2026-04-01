package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/songwaad/cs-event-backend/internal/handler"
)

func RegisterAuthRoutes(app *fiber.App, h *handler.UserHandler, secretKey string) {
	app.Post("/auth/login", func(c *fiber.Ctx) error {
		return h.Login(c, secretKey)
	})
	app.Post("/auth/logout", h.Logout)
	app.Post("/auth/register", h.Register)
}

func RegisterUserRoutes(app *fiber.App, h *handler.UserHandler) {
	app.Get("/me", h.GetMe)
	app.Get("/users", h.GetAllUsers)
	app.Get("/user/:id", h.GetUserByID)
	app.Delete("/user/:id", h.DeleteUser)
	app.Patch("/user/:id/change-password", h.ChangePassword)
	app.Patch("/user/:id/role", h.UpdateRole)
	app.Patch("/user/:id/status", h.UpdateStatus)
}
