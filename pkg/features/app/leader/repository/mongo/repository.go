package mongo

import (
	"github.com/amorindev/go-tmpl/pkg/features/app/leader/port"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// Make sure Repository implements ports.LeaderRepo
// at compile time
var _ port.LeaderRepo = &Repository{}

type Repository struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func NewLeaderRepo(client *mongo.Client, collection *mongo.Collection) *Repository {
	return &Repository{
		Client:     client,
		Collection: collection,
	}
}
