package main

import (
	"airdrop/internal/repositories"
	"airdrop/internal/users"
	"airdrop/pkg/storage/minio"
	"airdrop/pkg/storage/redis"
	"github.com/gofiber/fiber/v2"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Ldate | log.Lshortfile)

	log.Println(os.Getenv("DATABASE_URL"))
	app := fiber.New()
	db, err := sqlx.Connect("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	redisClient := redis.NewClient()
	minioClient := minio.NewClient()

	_ = redis.NewService(redisClient)
	_ = minio.NewService(minioClient)
	_ = repositories.NewTaskRepository(db)

	userRepo := repositories.NewUserRepository(db)

	users.SetupRoutes(app, userRepo)

	err = app.Listen(":5000")
	if err != nil {
		panic(err)
	}
}
