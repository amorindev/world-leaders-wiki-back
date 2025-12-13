package service

import "github.com/amorindev/go-tmpl/pkg/features/app/political-party/port"

var _ port.PPartySrv = &Service{}

type Service struct {
	PPartyRepo port.PPartyRepo
}

func NewPPartySrv(pPartyRepo port.PPartyRepo) *Service {
	return &Service{
        PPartyRepo: pPartyRepo,
    }
}
