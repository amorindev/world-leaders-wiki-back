package core

import (
	"net/url"
	"strings"
	"time"

	"github.com/amorindev/go-tmpl/pkg/shared/domain"
	sharedD "github.com/amorindev/go-tmpl/pkg/shared/domain"
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

	if err := validateSocialURL(req.Facebook, "facebook"); err != nil {
		return err
	}
	if err := validateSocialURL(req.Instagram, "instagram"); err != nil {
		return err
	}
	if err := validateOptionalURL(req.Twitter, "twitter"); err != nil {
		return err
	}
	if err := validateOptionalURL(req.YouTube, "youtube"); err != nil {
		return err
	}
	if err := validateOptionalURL(req.Linkedin, "linkedin"); err != nil {
		return err
	}
	if err := validateOptionalURL(req.Website, "website"); err != nil {
		return err
	}

	return nil
}

func validateSocialURL(urlStr, field string) error {
	if strings.TrimSpace(urlStr) == "" {
		return sharedD.NewAppError(domain.ErrCodeInvalidParams, field+" is required")
	}
	return validateURL(urlStr, field)
}

func validateOptionalURL(urlStr *string, field string) error {
	if urlStr == nil || strings.TrimSpace(*urlStr) == "" {
		return nil
	}
	return validateURL(*urlStr, field)
}

func validateURL(u, field string) error {
	parsed, err := url.ParseRequestURI(u)
	if err != nil || parsed.Scheme == "" || parsed.Host == "" {
		return sharedD.NewAppError(domain.ErrCodeInvalidParams, field+" must be a valid URL")
	}
	return nil
}
