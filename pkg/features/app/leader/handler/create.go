package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/amorindev/go-tmpl/pkg/features/app/leader/core"
	"github.com/amorindev/go-tmpl/pkg/features/app/leader/domain"
	sharedC "github.com/amorindev/go-tmpl/pkg/shared/api/core"
	sharedD "github.com/amorindev/go-tmpl/pkg/shared/domain"
)

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req core.CreateLeaderReq

	// Decode JSON request body into SignUpReq struct
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		sharedC.RespondError(w, sharedD.NewAppError(sharedD.ErrCodeInvalidParams, "invalid request body"))
		return
	}

	defer r.Body.Close()

	// Validate the sign-up request
	err = req.IsCreateLeaderValid()
	if err != nil {
		sharedC.RespondError(w, err)
		return
	}

	leader := &domain.Leader{
		FullName:    req.FullName,
		NickName:    req.Nickname,
		Phrase:      req.Phrase,
		Biography:   req.Biography,
		BirthDate:   req.BirthDate,
		Nationality: req.Nationality,
		Gender:      req.Gender,
		Ideology:    req.Ideology,
		Facebook:    req.Facebook,
		Instagram:   req.Instagram,
		Twitter:     req.Twitter,
		YouTube:     req.YouTube,
		Linkedin:    req.Linkedin,
		Website:     req.Website,
	}

	err = h.LeaderSrv.Create(context.Background(), leader)
	if err != nil {
		sharedC.RespondError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(leader)
}
