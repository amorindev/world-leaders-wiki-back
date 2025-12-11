package mongo

import (
	"context"
	"fmt"

	"github.com/amorindev/go-tmpl/pkg/shared/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) UpdateAvatarPath(ctx context.Context, leaderID, imgPath string) error {
	oID, err := bson.ObjectIDFromHex(leaderID)
	if err != nil {
		return fmt.Errorf("%w: failed to convert leaderID to ObjectID :%w", domain.ErrIncorrectID, err)
	}

	filter := bson.M{"_id": oID}
	update := bson.M{"$set": bson.M{"avatar_path": imgPath}}

	result, err := r.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update leader avatar path in database: %w", err)
	}

	if result.MatchedCount == 0 {
		return domain.ErrNotFound
	}

	return nil
}