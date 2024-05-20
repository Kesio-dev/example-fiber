package controllers

import (
	"airdrop/internal/models"
	"airdrop/internal/tasks/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type TaskController struct {
	service *services.TaskService
}

func NewTaskController(service *services.TaskService) *TaskController {
	return &TaskController{service: service}
}

func (c *TaskController) Get(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(400).SendString("Invalid ID")
	}
	user, err := c.service.GetByID(id)
	if err != nil {
		return ctx.Status(404).SendString("User Not Found")
	}
	return ctx.JSON(user)
}

func (c *TaskController) Create(ctx *fiber.Ctx) error {
	user := new(models.TaskModel)
	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(400).SendString(err.Error())
	}
	if err := c.service.Create(user); err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	return ctx.Status(201).JSON(user)
}
