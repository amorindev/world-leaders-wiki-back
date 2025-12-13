package handler

import (
	"net/http"

	"github.com/amorindev/go-tmpl/pkg/features/app/political-party/port"
)

type Handler struct {
	PPartySrv port.PPartySrv
}

func NewPPartyHandler(server *http.ServeMux, pPartySrv port.PPartySrv) *Handler {
	h := &Handler{
		PPartySrv: pPartySrv,
	}

	server.HandleFunc("POST /p-parties", h.Create)

	return h
}
