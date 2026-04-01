package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/songwaad/cs-event-backend/internal/handler"
)

func RegisterStrategyRoutes(app *fiber.App, h *handler.HttpStrategyHandler) {
	app.Get("/strategies", h.GetAllStrategies)
	app.Get("/strategy/:id", h.GetStrategyByID)
	app.Get("/event-strategies", h.GetAllEventStrategies)
	app.Get("/event-strategy/:id", h.GetEventStrategyByID)
}
