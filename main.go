package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	"github.com/songwaad/cs-event-backend/adapters"
	"github.com/songwaad/cs-event-backend/config"
	_ "github.com/songwaad/cs-event-backend/docs"
	"github.com/songwaad/cs-event-backend/entities"
	middleware "github.com/songwaad/cs-event-backend/middlewares"

	// middleware "github.com/songwaad/cs-event-backend/middlewares"
	"github.com/songwaad/cs-event-backend/usecases"
)

// @title ComSci Event API
// @description This is a sample server for a ComSci Event API.
// @version 1.0
// @host localhost:8080
// @BasePath /
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	secretKey := os.Getenv("JWT_SECRET_KEY")

	DB := config.GetDatabase()
	DB.AutoMigrate(&entities.User{})
	DB.AutoMigrate(&entities.Speaker{})
	DB.AutoMigrate(&entities.Event{})
	DB.AutoMigrate(&entities.Budget{})
	DB.AutoMigrate(&entities.Notification{})

	userRepo := adapters.NewGormUserRepo(DB)
	UserService := usecases.NewUserService(userRepo)
	userHandler := adapters.NewUserHandler(UserService)

	speakerRepo := adapters.NewGormSpeakerRepo(DB)
	speakerService := usecases.NewSpeakerService(speakerRepo)
	speakerHandler := adapters.NewHttpSpeakerHandler(speakerService)

	strategyRepo := adapters.NewGormStrategyRepo(DB)
	strategyService := usecases.NewStrategyService(strategyRepo)
	strategyHandler := adapters.NewHttpStrategyHandler(strategyService)

	notifyRepo := adapters.NewGormNotificationRepo(DB)
	notifyService := usecases.NewNotificationService(notifyRepo)
	notifyHandler := adapters.NewHttpNotificationHandler(notifyService)

	eventRepo := adapters.NewGormEventRepo(DB)
	eventService := usecases.NewEventService(eventRepo)
	eventHandler := adapters.NewHttpEventHandler(eventService, UserService, notifyService)

	budgetRepo := adapters.NewGormBudgetRepo(DB)
	budgetService := usecases.NewBudgetService(budgetRepo)
	budgetHandler := adapters.NewHttpBudgetHandler(budgetService)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowMethods:     "GET,POST,PUT,PATCH,DELETE",
		AllowHeaders:     "Content-Type, Authorization",
		AllowCredentials: true,
	}))

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Post("/auth/login", func(c *fiber.Ctx) error {
		return userHandler.Login(c, secretKey)
	})
	app.Post("/auth/logout", userHandler.Logout)
	app.Post("/auth/register", userHandler.Register)

	app.Use(middleware.AuthMiddleware())

	app.Get("/user/:id", userHandler.GetUserByID)
	app.Delete("/user/:id", userHandler.DeleteUser)
	app.Patch("/user/:id/change-password", userHandler.ChangePassword)
	app.Patch("/user/:id/role", userHandler.UpdateRole)
	app.Patch("/user/:id/status", userHandler.UpdateStatus)
	app.Get("/users", userHandler.GetAllUsers)

	app.Get("/budgets", budgetHandler.GetAllBudgets)
	app.Get("/budget/:id", budgetHandler.GetBudgetByID)
	app.Post("/budget", budgetHandler.CreateBudget)
	app.Patch("/budget/:id", budgetHandler.UpdateBudget)

	app.Get("/event-strategies", strategyHandler.GetAllEventStrategies)
	app.Get("/event-strategy/:id", strategyHandler.GetEventStrategyByID)

	app.Get("/events/plans", eventHandler.GetAllEventPlan)
	app.Get("/events/status", eventHandler.GetAllEventStatus)
	app.Get("/events/types", eventHandler.GetAllEventType)
	app.Get("/events/types-status", eventHandler.GetAllEventTypeStatus)

	app.Get("/events", eventHandler.GetAllEvents)
	app.Get("/event/:id", eventHandler.GetEventByID)
	app.Post("/event", eventHandler.CreateEvent)
	app.Put("/event/:id", eventHandler.UpdateEvent)
	app.Delete("/event/:id", eventHandler.DeleteEvent)

	app.Patch("/event/:id/status", eventHandler.UpdateEventStatus)
	app.Get("/calendar", eventHandler.GetCalendarEvents)

	app.Get("/speakers", speakerHandler.GetAllSpeakers)
	app.Get("/speaker/:id", speakerHandler.GetSpeakerByID)
	app.Post("/speaker", speakerHandler.CreateSpeaker)
	app.Patch("/speaker/:id", speakerHandler.UpdateSpeaker)
	app.Delete("/speaker/:id", speakerHandler.DeleteSpeaker)

	app.Get("/strategies", strategyHandler.GetAllStrategies)
	app.Get("/strategy/:id", strategyHandler.GetStrategyByID)

	app.Get("/notification/user/:id", notifyHandler.GetNotifyByID)
	app.Patch("/notification/:id/inactive", notifyHandler.InActive)
	app.Delete("/notification/:id", notifyHandler.DeleteNotify)

	app.Listen(":8080")
}
