package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/songwaad/cs-event-backend/internal/handler"
)

func RegisterBudgetRoutes(app *fiber.App, h *handler.HttpBudgetHandler) {
	app.Get("/budgets", h.GetAllBudgets)
	app.Get("/budget/:id", h.GetBudgetByID)
	app.Post("/budget", h.CreateBudget)
	app.Patch("/budget/:id", h.UpdateBudget)
	app.Delete("/budget/:id", h.DeleteBudget)
}
