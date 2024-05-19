package main

import (
	"airdrop/internal/tasks"
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

	redisService := redis.NewService(redisClient)
	_ = minio.NewService(minioClient)
	users.SetupRoutes(app, db)
	tasks.SetupRoutes(app, db, redisService)

	err = app.Listen(":5000")
	if err != nil {
		panic(err)
	}
}
