package repository

import (
	"github.com/Sup3r-Us3r/barber-server/internal/domain/gateway"
	"github.com/Sup3r-Us3r/barber-server/internal/infra/repository/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

type RepositoryContainer struct {
	BarberRepository gateway.BarberGatewayInterface
}

var RepositoryContainerInstance RepositoryContainer

func GetMongoDBRepositories(database *mongo.Database) *RepositoryContainer {
	repositoryContainer := &RepositoryContainer{
		BarberRepository: mongodb.NewBarberRepositoryMongo(database),
	}
	RepositoryContainerInstance = *repositoryContainer

	return repositoryContainer
}
