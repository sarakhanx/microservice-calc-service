package handlers

import (
	"microservices/internal/core/ports"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CalculatorHandler struct {
	service ports.CalculatorService
}

func NewCalculatorHandler(service ports.CalculatorService) *CalculatorHandler {
	return &CalculatorHandler{service: service}
}

func (h *CalculatorHandler) Add(c *fiber.Ctx) error {
	a, err := strconv.Atoi(c.Query("a"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid parameter a"})
	}
	b, err := strconv.Atoi(c.Query("b"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid parameter b"})
	}
	return c.JSON(fiber.Map{"result": h.service.Add(a, b)})
}

func (h *CalculatorHandler) Subtract(c *fiber.Ctx) error {
	a, err := strconv.Atoi(c.Query("a"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid parameter a"})
	}
	b, err := strconv.Atoi(c.Query("b"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid parameter b"})
	}
	return c.JSON(fiber.Map{"result": h.service.Subtract(a, b)})
}
