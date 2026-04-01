package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	_ "github.com/songwaad/cs-event-backend/docs"
	"github.com/songwaad/cs-event-backend/internal/config"
	"github.com/songwaad/cs-event-backend/internal/database"
	"github.com/songwaad/cs-event-backend/internal/handler"
	"github.com/songwaad/cs-event-backend/internal/repository"
	"github.com/songwaad/cs-event-backend/internal/routes"
	"github.com/songwaad/cs-event-backend/internal/usecase"
	"gorm.io/gorm"
)

// @title ComSci Event API
// @description This is a sample server for a ComSci Event API.
// @version 1.0
// @host localhost:3000
// @BasePath /
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cfg := config.Load()
	db := database.Connect(cfg.DB)
	database.Migrate(db)

	app := fiber.New()
	setupMiddleware(app)
	routes.Register(app, initHandlers(db), cfg.JWTSecret)
	app.Listen(":3000")
}

func setupMiddleware(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowMethods:     "GET,POST,PUT,PATCH,DELETE",
		AllowHeaders:     "Content-Type, Authorization",
		AllowCredentials: true,
	}))
	app.Get("/swagger/*", swagger.HandlerDefault)
}

func initHandlers(db *gorm.DB) routes.Handlers {
	userRepo := repository.NewGormUserRepo(db)
	userService := usecase.NewUserService(userRepo)

	speakerRepo := repository.NewGormSpeakerRepo(db)
	speakerService := usecase.NewSpeakerService(speakerRepo)

	strategyRepo := repository.NewGormStrategyRepo(db)
	strategyService := usecase.NewStrategyService(strategyRepo)

	notifyRepo := repository.NewGormNotificationRepo(db)
	notifyService := usecase.NewNotificationService(notifyRepo)

	eventRepo := repository.NewGormEventRepo(db)
	eventService := usecase.NewEventService(eventRepo)

	budgetRepo := repository.NewGormBudgetRepo(db)
	budgetService := usecase.NewBudgetService(budgetRepo)

	return routes.Handlers{
		User:     handler.NewUserHandler(userService),
		Speaker:  handler.NewHttpSpeakerHandler(speakerService),
		Strategy: handler.NewHttpStrategyHandler(strategyService),
		Notify:   handler.NewHttpNotificationHandler(notifyService),
		Event:    handler.NewHttpEventHandler(eventService, userService, notifyService),
		Budget:   handler.NewHttpBudgetHandler(budgetService),
	}
}
