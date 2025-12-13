package port

import (
	"context"

	"github.com/amorindev/go-tmpl/pkg/features/app/political-party/domain"
)

type PPartyRepo interface {
	Insert(ctx context.Context, pParty *domain.PoliticalParty) error
}

type PPartySrv interface {
	Create(ctx context.Context, pParty *domain.PoliticalParty) error
}
