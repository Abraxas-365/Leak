package user

import (
	"github.com/gofiber/fiber/v2"

	"github/abraxas-365/Leak/internal/mongo"
	app "github/abraxas-365/Leak/pkg/user/application"
	"github/abraxas-365/Leak/pkg/user/infrastructure/repository"
)

func ModuleFactory(appFiber *fiber.App, mongo mongo.MongoConnection) app.Application {
	repo := repository.RepositoryFactory(mongo, "users")
	application := app.AppFactory(repo)

	return application
}
