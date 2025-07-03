package handler

import (
	"hexagonal_cheese/models/input"
	"hexagonal_cheese/service"
	"net/http"
	"strconv"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type CheeseHandler struct {
	Cheese service.CheeseService
}

func NewCheeseHandler(s service.CheeseService) CheeseHandler {
	return CheeseHandler{Cheese: s}
}

func (h CheeseHandler) GetAllCheese(c *fiber.Ctx) error {
    fmt.Println("Handler reached")
    cheeses, err := h.Cheese.GetAllCheese()
    if err != nil {
        fmt.Println("Service error:", err)
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "message": "Failed to get cheeses error 500",
        })
    }
    return c.Status(http.StatusOK).JSON(fiber.Map{
        "data": cheeses,
    })
}

func (h CheeseHandler) GetCheese(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	cheese, err := h.Cheese.GetCheese(id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Cheese not found",
		})
	}
	return c.JSON(cheese)
}

func (h CheeseHandler) CreateCheese(c *fiber.Ctx) error {
	var in input.Cheese
	if err := c.BodyParser(&in); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := h.Cheese.CreateCheese(in); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create cheese"})
	}
	return c.SendStatus(fiber.StatusCreated)
}

func (h CheeseHandler) UpdateCheese(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var in input.Cheese
	if err := c.BodyParser(&in); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := h.Cheese.UpdateCheese(id, in); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update cheese"})
	}
	return c.SendStatus(fiber.StatusOK)
}

func (h CheeseHandler) DeleteCheese(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := h.Cheese.DeleteCheese(id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete cheese"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func (h CheeseHandler) GetMostCountryCheese(c *fiber.Ctx) error{
	cheese, magnitude , err := h.Cheese.GetMostCountryCheese()
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Country Not Found",
	})
	}
	return c.JSON(fiber.Map{
		"country":  cheese,
		"no_of_cheese": magnitude,
	})
}
