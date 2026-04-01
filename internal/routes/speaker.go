package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/songwaad/cs-event-backend/internal/handler"
)

func RegisterSpeakerRoutes(app *fiber.App, h *handler.HttpSpeakerHandler) {
	app.Get("/speakers", h.GetAllSpeakers)
	app.Get("/speaker/:id", h.GetSpeakerByID)
	app.Post("/speaker", h.CreateSpeaker)
	app.Patch("/speaker/:id", h.UpdateSpeaker)
	app.Delete("/speaker/:id", h.DeleteSpeaker)
}
