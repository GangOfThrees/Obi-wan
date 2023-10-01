package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/GangOfThrees/Obi-wan/darth-vader/internal/constants"
	"github.com/GangOfThrees/Obi-wan/darth-vader/internal/controllers"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	app := fiber.New()

	controllers.RegisterMiddlewares(app)
	controllers.RegisterRoutes(app)

	fmt.Println("Server starting")

	app.Listen(fmt.Sprintf(":%s", os.Getenv(constants.SERVER_PORT)))

	fmt.Printf("Server started successfully on port %s\n", os.Getenv(constants.SERVER_PORT))
}
