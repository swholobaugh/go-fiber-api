package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/swholobaugh/go-fiber/config"
	"github.com/swholobaugh/go-fiber/routes"
)

func main() {
	//Initialize fiber
	app := fiber.New()

	//Middleware
	app.Use(logger.New())

	// dotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	// config db
	config.ConnectDB()

	//Setup routes
	setupRoutes(app)

	//Listen on server 8000 and catch any errors
	err = app.Listen(":8000")

	//handle error
	if err != nil {
		panic(err)
	}
}

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the endpoint",
		})
	})

	//api group
	api := app.Group("/api")

	//give response when at /api
	api.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the api endpoint",
		})
	})

	//connect todo routes
	routes.TodoRoute(api.Group("/todos"))
}
