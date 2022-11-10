package waitinglist

import (
	"github/abraxas-365/Leak/internal/mongo"
	"github/abraxas-365/Leak/pkg/waiting-list/application"
	"github/abraxas-365/Leak/pkg/waiting-list/infrastructure/repository"
	"github/abraxas-365/Leak/pkg/waiting-list/infrastructure/rest"

	"github.com/gofiber/fiber/v2"
)

func ModuleFactory(appFiber *fiber.App, mongo mongo.MongoConnection) {
	repo := repository.RepositoryFactory(mongo, "waitinglist")
	application := application.AppFactory(repo)
	handler := rest.HandlerFabric(application)
	rest.Routes(appFiber, handler)
}
