package main

import (
	"log"
	"microservices/internal/core/services"
	"microservices/internal/handlers"
	"microservices/internal/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	calculatorService := services.NewCalculatorService()
	calculatorHandler := handlers.NewCalculatorHandler(calculatorService)

	route.SetupCalculatorRoutes(app, calculatorHandler)
	log.Fatal(app.Listen(":3003"))
}
