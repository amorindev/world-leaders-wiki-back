package core

import (
	"regexp"
	"strings"
	"time"

	"github.com/amorindev/go-tmpl/pkg/shared/domain"
	sharedD "github.com/amorindev/go-tmpl/pkg/shared/domain"
	"github.com/amorindev/go-tmpl/pkg/shared/validator"
)

// CreateLeaderReq represents the request structure for Leader creation
type CreateLeaderReq struct {
	Slug      string  `json:"slug"`
	FullName  string  `json:"full_name"`
	Phrase    *string `json:"phrase"`
	Nickname  *string `json:"nickname"`
	Biography string  `json:"biography"`

	BirthDate   *time.Time `json:"birth_date"`
	DeathDate   *time.Time `json:"death_date"`
	Nationality string     `json:"nationality"`
	Gender      string     `json:"gender"`
	Ideology    *string    `json:"ideology"`

	Facebook  *string `json:"facebook"`
	Instagram *string `json:"instagram"`
	Twitter   *string `json:"twitter"`
	YouTube   *string `json:"youtube"`
	Linkedin  *string `json:"linkedin"`
	Website   *string `json:"website"`
}

// slugRegex validates SEO-friendly slugs.
// Accepted:
//   - lowercase letters (a–z)
//   - numbers (0–9)
//   - single hyphens between words
//
// Examples (valid):
//
//	rafael
//	rafael-lopez
//	rafael-lopez-aliaga
//
// Examples (invalid):
//
//	Rafael-Lopez   // uppercase letters
//	rafael_lopez   // underscores (_)
//	rafael lopez   // spaces
//	rafael--lopez  // double hyphens
//	-rafael        // starts with hyphen
//	rafael-        // ends with hyphen
//	rafaél         // accented characters
//	rafael!        // special characters
//	@rafael        // '@' is not allowed
var slugRegex = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)

func (req *CreateLeaderReq) IsCreateLeaderValid() error {
	// Validate email field is not empty
	if strings.TrimSpace(req.FullName) == "" {
		return sharedD.NewAppError(domain.ErrCodeInvalidParams, "full_name is required")
	}

	if strings.TrimSpace(req.Slug) == "" {
		return sharedD.NewAppError(domain.ErrCodeInvalidParams, "slug is required")
	}

	if len(req.Slug) > 80 {
		return sharedD.NewAppError(domain.ErrCodeInvalidParams, "slug is too long")
	}

	if !slugRegex.MatchString(req.Slug) {
		return sharedD.NewAppError(
			domain.ErrCodeInvalidParams,
			"slug format is invalid (use lowercase letters, numbers and hyphens)",
		)
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

	if req.DeathDate != nil {
		// Prevent future death dates
		if req.DeathDate.After(time.Now()) {
			return sharedD.NewAppError(
				domain.ErrCodeInvalidParams,
				"death_date cannot be in the future",
			)
		}

		// Death date must be after birth date
		if req.BirthDate != nil && req.DeathDate.Before(*req.BirthDate) {
			return sharedD.NewAppError(
				domain.ErrCodeInvalidParams,
				"death_date cannot be before birth_date",
			)
		}
	}

	if err := validator.ValidateOptionalURL(req.Facebook, "facebook"); err != nil {
		return sharedD.NewAppError(domain.ErrCodeInvalidParams, err.Error())
	}
	if err := validator.ValidateOptionalURL(req.Instagram, "instagram"); err != nil {
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
