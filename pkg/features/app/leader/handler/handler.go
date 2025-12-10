package handler

import (
	"net/http"

	"github.com/amorindev/go-tmpl/pkg/features/app/leader/port"
)

type Handler struct {
	LeaderSrv port.LeaderSrv
}

func NewLeaderHandler(server *http.ServeMux, leaderSrv port.LeaderSrv) *Handler {
	h := &Handler{
		LeaderSrv: leaderSrv,
	}

	server.HandleFunc("POST /leaders", h.Create)
	
	return h
}