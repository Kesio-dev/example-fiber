package users

import (
	"airdrop/internal/users/controllers"
	"airdrop/internal/users/repositories"
	"airdrop/internal/users/services"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func SetupRoutes(app *fiber.App, db *sqlx.DB) {
	// Создаем репозиторий
	userRepo := repositories.NewUserRepository(db)

	// Создаем сервис
	userService := services.NewUserService(userRepo)

	// Создаем контроллер
	userController := controllers.NewUserController(userService)

	// Настраиваем маршруты
	app.Get("/users/:id", userController.GetUser)
	app.Post("/users", userController.CreateUser)
}
