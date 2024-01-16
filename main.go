package main

import (
	"context"
	"fmt"
	"github.com/dev-hyunsang/clone-stackbuck-backend/db"
	"github.com/dev-hyunsang/clone-stackbuck-backend/middleware"
	"github.com/dev-hyunsang/clone-stackbuck-backend/models"
	"github.com/gofiber/fiber/v2"
	"log"
)

var (
	port = ":3000"
)

func main() {
	app := fiber.New()

	client, err := db.ConnectMySQL()
	if err != nil {
		log.Fatalln(err)
	}

	if err := client.AutoMigrate(models.Users{}); err != nil {
		log.Fatalln(err)
	}

	redisClient, err := db.ConnectRedis()
	if err != nil {
		log.Fatalln(err)
	}

	result := redisClient.Ping(context.Background())

	middleware.Middleware(app)

	log.Println("[OK] Successfully Connect to MySQL")
	log.Println("[OK] Successfully Connect to Redis")
	log.Println("[OK] Successfully Ping to Redis\n", result)
	log.Println("[OK] Successfully AutoMigration in MySQL - Create Table")
	log.Println(fmt.Sprintf("[OK] Running on localhost%s", port))

	if err := app.Listen(port); err != nil {
		log.Fatalln(err)
	}
}
