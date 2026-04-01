package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/songwaad/cs-event-backend/internal/handler"
)

func RegisterEventRoutes(app *fiber.App, h *handler.HttpEventHandler) {
	app.Get("/events", h.GetAllEvents)
	app.Get("/events/plans", h.GetAllEventPlan)
	app.Get("/events/status", h.GetAllEventStatus)
	app.Get("/events/types", h.GetAllEventType)
	app.Get("/events/types-status", h.GetAllEventTypeStatus)
	app.Get("/event/:id", h.GetEventByID)
	app.Post("/event", h.CreateEvent)
	app.Put("/event/:id", h.UpdateEvent)
	app.Delete("/event/:id", h.DeleteEvent)
	app.Patch("/event/:id/status", h.UpdateEventStatus)
	app.Get("/calendar", h.GetCalendarEvents)
}
