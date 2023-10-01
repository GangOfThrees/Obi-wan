package controllers

import (
	"time"

	"github.com/GangOfThrees/Obi-wan/kenobi/internal/constants"
	"github.com/GangOfThrees/Obi-wan/kenobi/internal/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
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

	// use CORS
	// TODO: configure CORS
	app.Use(cors.New())

	// TODO: consider using compress middleware

	// use limiter middleware
	app.Use(limiter.New(
		limiter.Config{
			Max: 100,
		},
	))

	// use custom middlewares
	app.Use(middlewares.ValidateApiKeyMiddleware)
}

func RegisterRoutes(app *fiber.App) {
	// register bot routes
	RegisterBotRoutes(app)

	// TODO: register health check route
}

func RegisterBotRoutes(app *fiber.App) {
	botRoutes := app.Group(constants.ENDPOINT_BOT)
	botRoutes.Post(constants.ENDPOINT_KNOWLEDGE_BASE, KnowledgeBase)
}
