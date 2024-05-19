package controllers

import (
	"airdrop/internal/users/models"
	"airdrop/internal/users/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) GetUser(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(400).SendString("Invalid ID")
	}
	user, err := c.service.GetUserByID(id)
	if err != nil {
		return ctx.Status(404).SendString("User Not Found")
	}
	return ctx.JSON(user)
}

func (c *UserController) CreateUser(ctx *fiber.Ctx) error {
	user := new(models.User)
	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(400).SendString(err.Error())
	}
	if err := c.service.CreateUser(user); err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	return ctx.Status(201).JSON(user)
}
