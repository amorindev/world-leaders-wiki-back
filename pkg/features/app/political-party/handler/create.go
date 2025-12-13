package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/amorindev/go-tmpl/pkg/features/app/political-party/core"
	"github.com/amorindev/go-tmpl/pkg/features/app/political-party/domain"
	sharedC "github.com/amorindev/go-tmpl/pkg/shared/api/core"
	sharedD "github.com/amorindev/go-tmpl/pkg/shared/domain"
)

// Create handles the HTTP request to create a new political party.
// It parses and validates the incoming JSON payload, maps it to the domain model,
// and delegates the creation process to the political party service.
func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req core.CreatePoliticalPartyReq

	// Decode JSON request body into SignUpReq struct
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		sharedC.RespondError(w, sharedD.NewAppError(sharedD.ErrCodeInvalidParams, "invalid request body"))
		return
	}

	defer r.Body.Close()

	// Validate the sign-up request
	err = req.IsCreatePoliticalPartyValid()
	if err != nil {
		sharedC.RespondError(w, err)
		return
	}

	pParty := &domain.PoliticalParty{
		Name:        req.Name,
		Acronym:     req.Acronym,
		Country:     req.Country,
		Description: req.Description,
		Founders:    req.Founders,
		Website:     req.Website,
		FoundedAt:   req.FoundedAt,
	}

	err = h.PPartySrv.Create(context.Background(), pParty)
	if err != nil {
		sharedC.RespondError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pParty)
}
