package service

import (
	"context"
	"time"

	"github.com/amorindev/go-tmpl/pkg/features/app/political-party/domain"
	sharedD "github.com/amorindev/go-tmpl/pkg/shared/domain"
)

func (s *Service) Create(ctx context.Context, pParty *domain.PoliticalParty) error {
	now := time.Now().UTC()
	pParty.CreatedAt = &now

	err := s.PPartyRepo.Insert(ctx, pParty)
	if err != nil {
		return sharedD.ManageError(err, "error creating political party")
	}

	return nil
}
