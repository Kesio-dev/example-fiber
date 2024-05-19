package users

import (
	"airdrop/internal/repositories"
	"airdrop/internal/users/controllers"
	"airdrop/internal/users/services"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, userRepo *repositories.UserRepository) {
	userService := services.NewUserService(userRepo)

	// Создаем контроллер
	userController := controllers.NewUserController(userService)

	// Настраиваем маршруты
	app.Get("/users", userController.GetUsers)
	app.Get("/users/:id", userController.GetUser)
	app.Post("/users", userController.CreateUser)
}
