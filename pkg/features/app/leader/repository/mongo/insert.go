package mongo

import (
	"context"
	"fmt"

	"github.com/amorindev/go-tmpl/pkg/features/app/leader/domain"
	sharedD "github.com/amorindev/go-tmpl/pkg/shared/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) Insert(ctx context.Context, leader *domain.Leader) error {
	id := bson.NewObjectID()
	leader.ID = id

	_, err := r.Collection.InsertOne(ctx, leader)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return fmt.Errorf("%w: error inserting leader: %s", sharedD.ErrDuplicateKey, err.Error())
		}
		return fmt.Errorf("error inserting leader: %s", err.Error())
	}
	leader.ID = id.Hex()

	return nil
}
