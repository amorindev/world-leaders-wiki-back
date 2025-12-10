package service

import (
	"context"
	"time"

	sharedD "github.com/amorindev/go-tmpl/pkg/shared/domain"
	"github.com/amorindev/go-tmpl/pkg/features/app/leader/domain"
)

func (s *Service) Create(ctx context.Context, leader *domain.Leader) error {
	now := time.Now().UTC()
	leader.CreatedAt = &now

	err := s.LeaderRepo.Insert(ctx, leader)
	if err != nil {
	  return sharedD.ManageError(err, "error creating leader")
	}

	return nil
}
