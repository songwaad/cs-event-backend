package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/songwaad/cs-event-backend/internal/handler"
	middleware "github.com/songwaad/cs-event-backend/internal/middleware"
)

type Handlers struct {
	User     *handler.UserHandler
	Speaker  *handler.HttpSpeakerHandler
	Strategy *handler.HttpStrategyHandler
	Notify   *handler.HttpNotificationHandler
	Event    *handler.HttpEventHandler
	Budget   *handler.HttpBudgetHandler
}

func Register(app *fiber.App, h Handlers, secretKey string) {
	RegisterAuthRoutes(app, h.User, secretKey)

	app.Use(middleware.AuthMiddleware())

	RegisterUserRoutes(app, h.User)
	RegisterEventRoutes(app, h.Event)
	RegisterSpeakerRoutes(app, h.Speaker)
	RegisterBudgetRoutes(app, h.Budget)
	RegisterStrategyRoutes(app, h.Strategy)
	RegisterNotificationRoutes(app, h.Notify)
}
