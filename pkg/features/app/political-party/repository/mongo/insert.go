package mongo

import (
	"context"
	"fmt"

	"github.com/amorindev/go-tmpl/pkg/features/app/political-party/domain"
	sharedD "github.com/amorindev/go-tmpl/pkg/shared/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) Insert(ctx context.Context, pParty *domain.PoliticalParty) error {
	id := bson.NewObjectID()
	pParty.ID = id

	var oIDs []interface{}
	for _, v := range pParty.Founders {
		oID, err := bson.ObjectIDFromHex(v.(string))
		if err != nil {

		}
		oIDs = append(oIDs, oID)
	}

	pParty.Founders = oIDs

	_, err := r.Collection.InsertOne(ctx, pParty)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return fmt.Errorf("%w: error inserting political party: %s", sharedD.ErrDuplicateKey, err.Error())
		}
		return fmt.Errorf("error inserting political party: %s", err.Error())
	}
	pParty.ID = id.Hex()

	return nil
}
