package handler

import (
	"encoding/json"
	"net/http"

	sharedC "github.com/amorindev/go-tmpl/pkg/shared/api/core"
)

// GetRandom handles the HTTP request to retrieve a random leader.
// It calls the leader service to select a leader at random
// and returns it as a JSON response with status 200 OK.
// If an error occurs, it responds with the corresponding error.
func (h Handler) GetRandom(w http.ResponseWriter, r *http.Request){
    leader, err := h.LeaderSrv.GetRandom(r.Context())
	if err != nil {
		sharedC.RespondError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(leader)
}