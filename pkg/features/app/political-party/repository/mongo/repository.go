package mongo

import (
	"github.com/amorindev/go-tmpl/pkg/features/app/political-party/port"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// Make sure Repository implements ports.PPartyRepo
// at compile time
var _ port.PPartyRepo = &Repository{}

type Repository struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func NewPPartyRepo(client *mongo.Client, collection *mongo.Collection) *Repository {
	return &Repository{
		Client:     client,
		Collection: collection,
	}
}