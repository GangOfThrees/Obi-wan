package controllers

import (
	"os"
	"time"

	"github.com/GangOfThrees/Obi-wan/darth-vader/internal/constants"
	"github.com/GangOfThrees/Obi-wan/darth-vader/internal/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/proxy"
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
	registerKenobiRoutes(app)

	// TODO: register health check route
}

func registerKenobiRoutes(app *fiber.App) {
	proxyConfig := proxy.Config{
		Timeout: 30 * time.Second, // TODO: consider extracting to env var
		Servers: []string{
			os.Getenv(constants.KENOBI_BASE_URL), // fwd to kenobi
		},
		ModifyRequest: func(c *fiber.Ctx) error {
			c.Request().Header.Add("X-Real-IP", c.IP())
			return nil
		},
		ModifyResponse: func(c *fiber.Ctx) error {
			c.Response().Header.Del(fiber.HeaderServer)
			return nil
		},
	}

	app.Group(constants.ENDPOINT_BOT, proxy.Balancer(proxyConfig))
	app.Group(constants.ENPOINT_USERS, proxy.Balancer(proxyConfig))
}
