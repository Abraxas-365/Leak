package repository

import (
	mclient "github/abraxas-365/Leak/internal/mongo"
	"github/abraxas-365/Leak/pkg/user/domain/ports"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	client     *mongo.Client
	database   string
	collection string
	timeout    time.Duration
}

func RepositoryFactory(mongoConnection mclient.MongoConnection, collection string) ports.Repo {
	return &repo{
		client:     mongoConnection.Client,
		collection: collection,
		database:   mongoConnection.Database,
		timeout:    mongoConnection.Timeout,
	}

}
