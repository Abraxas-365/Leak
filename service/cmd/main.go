package main

import (
	"github/abraxas-365/Leak/internal/mongo"
	"github/abraxas-365/Leak/pkg/auth"
	"github/abraxas-365/Leak/pkg/user"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	mongoUri := os.Getenv("MONGODB_URI")

	mongoClient, _ := mongo.MongoFactory(mongoUri, "Leak", 10)

	userModule := user.ModuleFactory(app, mongoClient)
	auth.ModuleFactory(app, userModule)

	app.Listen(":3000")

}
