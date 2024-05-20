package tasks

import (
	"airdrop/internal/repositories"
	"airdrop/internal/tasks/controllers"
	"airdrop/internal/tasks/services"
	"airdrop/pkg/storage"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, taskRepo *repositories.TaskRepository, redisClient storage.RedisClient) {
	taskService := services.NewTaskService(taskRepo, redisClient)
	taskController := controllers.NewTaskController(taskService)

	api := app.Group("tasks")
	api.Get("/:id", taskController.Get)
	api.Post("/", taskController.Create)
}
