package service

import "github.com/amorindev/go-tmpl/pkg/features/app/leader/port"

var _ port.LeaderSrv = &Service{}

type Service struct {
	LeaderRepo port.LeaderRepo
}

func NewLeaderSrv(leaderRepo port.LeaderRepo) *Service {
	return &Service{
		LeaderRepo: leaderRepo,
	}
}
