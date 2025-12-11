package service

import (
	"context"
	"fmt"
	"io"
	"path/filepath"

	"github.com/amorindev/go-tmpl/pkg/features/app/leader/domain"
	sharedD "github.com/amorindev/go-tmpl/pkg/shared/domain"
)

func (s *Service) UploadLeaderImage(ctx context.Context, leaderID, imgType, imgName string, file io.Reader, contentType string) error {
	exists, err := s.LeaderRepo.Exists(ctx, leaderID)
	if err != nil {
		return sharedD.ManageError(err, "")
	}

	if !exists {
		return sharedD.ManageError(sharedD.ErrNotFound, "")
	}

	ext := filepath.Ext(imgName)
	path := fmt.Sprintf("leaders/%s/%s%s", imgType, leaderID, ext)

	err = s.LeaderFileStg.UploadImage(ctx, path, file, contentType)
	if err != nil {
		return sharedD.ManageError(err, "")
	}

	// update in the database
	switch imgType {
	case domain.LeaderImageAvatar:
		err = s.LeaderRepo.UpdateAvatarPath(ctx, leaderID, path)
	case domain.LeaderImageBanner:
		err = s.LeaderRepo.UpdateBannerPath(ctx, leaderID, path)
	}

	if err != nil {
		return sharedD.ManageError(err, "")
	}

	return nil
}
