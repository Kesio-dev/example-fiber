package users

import (
	"airdrop/internal/repositories"
	"airdrop/internal/users/controllers"
	"airdrop/internal/users/services"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, userRepo *repositories.UserRepository) {
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	api := app.Group("/users")

	api.Get("/", userController.GetAll)
}
