package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var ConvoSessionStore = session.New(
	session.Config{
		KeyLookup:  "cookie:convo-id",
		Expiration: 1 * time.Hour,
	},
)

func RegisterMiddlewares(app *fiber.App) {
	// use logger middleware
	app.Use(logger.New())

	// use helmet middleware
	app.Use(helmet.New())

	// TODO: consider using compress middleware

	// TODO: consider using rate limiter despite darth-vader already using it
}

func RegisterRoutes(app *fiber.App) {
	registerBotRoutes(app)
	registerUserRoutes(app)

	// TODO: register health check route
}

func registerBotRoutes(app *fiber.App) {
	botRoutes := app.Group("/bot")
	botRoutes.Post("/kb", KnowledgeBase)
}

func registerUserRoutes(app *fiber.App) {
	userRoutes := app.Group("/users")
	userRoutes.Post("/", CreateUser)
}
