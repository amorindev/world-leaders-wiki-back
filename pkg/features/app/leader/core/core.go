package core

import (
	"strings"
	"time"

	"github.com/amorindev/go-tmpl/pkg/shared/domain"
	sharedD "github.com/amorindev/go-tmpl/pkg/shared/domain"
	"github.com/amorindev/go-tmpl/pkg/shared/validator"
)

// CreateLeaderReq represents the request structure for Leader creation
type CreateLeaderReq struct {
	FullName  string  `json:"full_name"`
	Phrase    *string `json:"phrase"`
	Nickname  *string `json:"nickname"`
	Biography string  `json:"biography"`

	BirthDate   *time.Time `json:"birth_date"`
	Nationality string     `json:"nationality"`
	Gender      string     `json:"gender"`
	Ideology    *string    `json:"ideology"`

	Facebook  string  `json:"facebook"`
	Instagram string  `json:"instagram"`
	Twitter   *string `json:"twitter"`
	YouTube   *string `json:"youtube"`
	Linkedin  *string `json:"linkedin"`
	Website   *string `json:"website"`
}

func (req *CreateLeaderReq) IsCreateLeaderValid() error {
	// Validate email field is not empty
	if strings.TrimSpace(req.FullName) == "" {
		return sharedD.NewAppError(domain.ErrCodeInvalidParams, "full_name is required")
	}

	if strings.TrimSpace(req.Biography) == "" {
		return sharedD.NewAppError(domain.ErrCodeInvalidParams, "biography is required")
	}

	if req.Phrase != nil && strings.TrimSpace(*req.Phrase) == "" {
		return sharedD.NewAppError(domain.ErrCodeInvalidParams, "phrase is empty")
	}

	if req.Ideology != nil && strings.TrimSpace(*req.Ideology) == "" {
		return sharedD.NewAppError(domain.ErrCodeInvalidParams, "phrase is empty")
	}

	validGenders := map[string]bool{
		"male": true, "female": true, "other": true,
	}
	if !validGenders[strings.ToLower(req.Gender)] {
		return sharedD.NewAppError(domain.ErrCodeInvalidParams, "gender must be male, female or other")
	}

	if strings.TrimSpace(req.Nationality) == "" {
		return sharedD.NewAppError(domain.ErrCodeInvalidParams, "nationality is required")
	}

	if req.BirthDate != nil {
		// Prevent future birth dates
		if req.BirthDate.After(time.Now()) {
			return sharedD.NewAppError(domain.ErrCodeInvalidParams, "birth_date cannot be in the future")
		}
	}

	if err := validator.ValidateSocialURL(req.Facebook, "facebook"); err != nil {
		return sharedD.NewAppError(domain.ErrCodeInvalidParams, err.Error())
	}
	if err := validator.ValidateSocialURL(req.Instagram, "instagram"); err != nil {
		return sharedD.NewAppError(domain.ErrCodeInvalidParams, err.Error())
	}
	if err := validator.ValidateOptionalURL(req.Twitter, "twitter"); err != nil {
		return sharedD.NewAppError(domain.ErrCodeInvalidParams, err.Error())
	}
	if err := validator.ValidateOptionalURL(req.YouTube, "youtube"); err != nil {
		return sharedD.NewAppError(domain.ErrCodeInvalidParams, err.Error())
	}
	if err := validator.ValidateOptionalURL(req.Linkedin, "linkedin"); err != nil {
		return sharedD.NewAppError(domain.ErrCodeInvalidParams, err.Error())
	}
	if err := validator.ValidateOptionalURL(req.Website, "website"); err != nil {
		return sharedD.NewAppError(domain.ErrCodeInvalidParams, err.Error())
	}

	return nil
}
