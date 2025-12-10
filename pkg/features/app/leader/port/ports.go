package port

import (
	"context"

	"github.com/amorindev/go-tmpl/pkg/features/app/leader/domain"
)

type LeaderRepo interface {
	Insert(ctx context.Context, leader *domain.Leader) error
}

type LeaderSrv interface {
	Create(ctx context.Context, leader *domain.Leader) error
}
