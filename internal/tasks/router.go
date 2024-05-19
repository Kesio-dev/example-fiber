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
	taskService := services.NewTaskService(taskRepo, redisClient)
	taskController := controllers.NewTaskController(taskService)

	api := app.Group("tasks")
	api.Get("/:id", taskController.Get)
	api.Post("/", taskController.Create)
}
