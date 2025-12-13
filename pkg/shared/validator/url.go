package validator

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

func ValidateSocialURL(urlStr, field string) error {
	if strings.TrimSpace(urlStr) == "" {
		return fmt.Errorf("%s is required", field)
	}
	return validateURL(urlStr, field)
}

func ValidateOptionalURL(urlStr *string, field string) error {
	if urlStr == nil {
		return nil
	}
	if strings.TrimSpace(*urlStr) == "" {
		return errors.New("url must not be empty")
	}

	return validateURL(*urlStr, field)
}

func validateURL(u, field string) error {
	parsed, err := url.ParseRequestURI(u)
	if err != nil || parsed.Scheme == "" || parsed.Host == "" {
		return fmt.Errorf("%s must be a valid URL", field)
	}
	return nil
}
