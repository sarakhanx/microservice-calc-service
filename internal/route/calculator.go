package route

import (
	"microservices/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupCalculatorRoutes(app *fiber.App, calculatorHandler *handlers.CalculatorHandler) {
	calculator := app.Group("/calculator")

	calculator.Get("/add", calculatorHandler.Add)
	calculator.Get("/subtract", calculatorHandler.Subtract)
}
