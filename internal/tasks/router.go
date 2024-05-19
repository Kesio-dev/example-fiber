package tasks

import (
	"airdrop/internal/tasks/controllers"
	"airdrop/internal/tasks/repositories"
	"airdrop/internal/tasks/services"
	"airdrop/pkg/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func SetupRoutes(app *fiber.App, db *sqlx.DB, redisClient storage.RedisClient) {
	taskRepo := repositories.NewTaskRepository(db)
	userService := services.NewTaskService(taskRepo, redisClient)
	userController := controllers.NewUserController(userService)

	api := app.Group("tasks")
	api.Get("/:id", userController.Get)
	api.Post("/", userController.Create)
}
