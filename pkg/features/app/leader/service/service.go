package service

import "github.com/amorindev/go-tmpl/pkg/features/app/leader/port"

var _ port.LeaderSrv = &Service{}

type Service struct {
	LeaderRepo    port.LeaderRepo
	LeaderFileStg port.LeaderFileStg
}

func NewLeaderSrv(leaderRepo port.LeaderRepo, leaderFileStg port.LeaderFileStg) *Service {
	return &Service{
		LeaderRepo: leaderRepo,
		LeaderFileStg: leaderFileStg,
	}
}
