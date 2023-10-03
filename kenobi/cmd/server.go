package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/GangOfThrees/Obi-wan/kenobi/internal/constants"
	"github.com/GangOfThrees/Obi-wan/kenobi/internal/controllers"
	"github.com/GangOfThrees/Obi-wan/kenobi/internal/repository"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	app := fiber.New()

	err := repository.SetupDatabase()
	if err != nil {
		log.Fatal(err)
	}

	controllers.RegisterMiddlewares(app)
	controllers.RegisterRoutes(app)

	app.Listen(fmt.Sprintf(":%s", os.Getenv(constants.SERVER_PORT)))
	fmt.Printf("Server started successfully on port %s\n", os.Getenv(constants.SERVER_PORT))
}
