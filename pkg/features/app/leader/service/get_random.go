package service

import (
	"context"

	"github.com/amorindev/go-tmpl/pkg/features/app/leader/domain"
	sharedD "github.com/amorindev/go-tmpl/pkg/shared/domain"
)

func (s *Service) GetRandom(ctx context.Context) (*domain.Leader, error) {
	leader, err := s.LeaderRepo.GetRandom(ctx)
	if err != nil {
		return nil, sharedD.ManageError(err, "error getting random leader")
	}

	if leader.BannerPath != nil {
		url, err := s.LeaderFileStg.GetImage(ctx, *leader.BannerPath)
		if err != nil {
			return nil, sharedD.ManageError(err, "error getting leader banner")
		}
		leader.BannerUrl = &url
	}

	if leader.AvatarPath != nil {
		url, err := s.LeaderFileStg.GetImage(ctx, *leader.AvatarPath)
		if err != nil {
			return nil, sharedD.ManageError(err, "error getting leader avatar")
		}
		leader.AvatarUrl = &url
	}

	return leader, err
}
