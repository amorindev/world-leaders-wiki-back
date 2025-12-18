package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// CreateIndexes sets a unique index on the "slug" field.
func (r *Repository) CreateIndexes() error {
	_, err := r.Collection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "slug", Value: 1}},
		Options: options.Index().SetUnique(true).SetName("slug"),
	})
	if err != nil {
		return fmt.Errorf("error creating slug index: %w", err)
	}
	return nil
}
