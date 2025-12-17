package mongo

import (
	"context"
	"fmt"

	"github.com/amorindev/go-tmpl/pkg/features/app/leader/domain"
	sharedD "github.com/amorindev/go-tmpl/pkg/shared/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) GetRandom(ctx context.Context) (*domain.Leader, error) {
	pipeline := mongo.Pipeline{
		{{Key: "$sample", Value: bson.D{{Key: "size", Value: 1}}}},
	}

	cursor, err := r.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("failed to aggregate random leader: %w", err)

	}
	defer cursor.Close(ctx)

	var leaders []domain.Leader
	if err := cursor.All(ctx, &leaders); err != nil {
		return nil, fmt.Errorf("failed to decode leader result: %w", err)
	}

	if len(leaders) == 0 {
		return nil, fmt.Errorf("%w: random leader not found", sharedD.ErrNotFound)
	}

	leader := leaders[0]

	// Parse MongoDB ObjectID to string
	if oid, ok := leader.ID.(bson.ObjectID); ok {
		leader.ID = oid.Hex()
	}

	return &leader, nil
}
