package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/songwaad/cs-event-backend/internal/handler"
)

func RegisterNotificationRoutes(app *fiber.App, h *handler.HttpNotificationHandler) {
	app.Get("/notification/user/:id", h.GetNotifyByID)
	app.Patch("/notification/:id/inactive", h.InActive)
	app.Delete("/notification/:id", h.DeleteNotify)
}
