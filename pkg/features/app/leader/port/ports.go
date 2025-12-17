package port

import (
	"context"
	"io"

	"github.com/amorindev/go-tmpl/pkg/features/app/leader/domain"
)

type LeaderRepo interface {
	Insert(ctx context.Context, leader *domain.Leader) error
	Exists(ctx context.Context, id string) (bool, error)
	UpdateAvatarPath(ctx context.Context, leaderID, imgPath string) error
	UpdateBannerPath(ctx context.Context, leaderID, imgPath string) error
	GetRandom(ctx context.Context) (*domain.Leader, error)
}

type LeaderSrv interface {
	Create(ctx context.Context, leader *domain.Leader) error
	UploadLeaderImage(ctx context.Context, leaderID, imgType, imgName string, file io.Reader, contentType string) error
	GetRandom(ctx context.Context) (*domain.Leader, error)
}

type LeaderFileStg interface {
	UploadImage(ctx context.Context, imgPath string, file io.Reader, contentType string) error
	GetImage(ctx context.Context, imgPath string) (string, error) 
}
