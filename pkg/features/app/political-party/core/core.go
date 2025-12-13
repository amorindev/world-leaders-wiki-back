package core

import (
	"fmt"
	"strings"
	"time"

	"github.com/amorindev/go-tmpl/pkg/shared/domain"
	"github.com/amorindev/go-tmpl/pkg/shared/validator"
)

type CreatePoliticalPartyReq struct {
	Name        string        `json:"name"`
	Acronym     *string       `json:"acronym"`
	Country     string        `json:"country"`
	Description string        `json:"description"`
	Website     *string       `json:"website"`
	Founders    []interface{} `json:"founders"`
	FoundedAt   *time.Time    `json:"founded_at"`
}

func (req *CreatePoliticalPartyReq) IsCreatePoliticalPartyValid() error {
	// Validate Name
	if strings.TrimSpace(req.Name) == "" {
		return domain.NewAppError(domain.ErrCodeInvalidParams, "name is required")
	}
	if len(req.Name) < 3 {
		return domain.NewAppError(domain.ErrCodeInvalidParams, "name must have at least 3 characters")
	}

	// Validate Acronym
	if req.Acronym != nil {
		if *req.Acronym == "" {
			return domain.NewAppError(domain.ErrCodeInvalidParams, "acronym must not be empty")
		}
		if len(*req.Acronym) > 10 {
			return domain.NewAppError(domain.ErrCodeInvalidParams, "acronym must not exceed 10 characters")
		}
	}

	// Validate Country
	if strings.TrimSpace(req.Country) == "" {
		return domain.NewAppError(domain.ErrCodeInvalidParams, "country is required")
	}

	// Validate Description
	if strings.TrimSpace(req.Description) == "" {
		return domain.NewAppError(domain.ErrCodeInvalidParams, "description is required")
	}

	// Validate Founders
	if len(req.Founders) == 0 {
		return domain.NewAppError(domain.ErrCodeInvalidParams, "at least one founder is required")
	}

	for i, f := range req.Founders {
		if f == nil {
			msg := fmt.Sprintf("founder at index %d cannot be null", i)
			return domain.NewAppError(domain.ErrCodeInvalidParams, msg)
		}
	}

	// Validate FoundedAt (required)
	if req.FoundedAt == nil {
		return domain.NewAppError(domain.ErrCodeInvalidParams, "founded_at is required")
	}

	if req.FoundedAt.After(time.Now()) {
		return domain.NewAppError(domain.ErrCodeInvalidParams, "founded_at cannot be a future date")
	}

	// Validate Website (optional)
	if err := validator.ValidateOptionalURL(req.Website, "website"); err != nil {
		return domain.NewAppError(domain.ErrCodeInvalidParams, err.Error())
	}

	return nil
}
