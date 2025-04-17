package config

import (
	"GoSamples/ServiceSample/internal/pkg/config/types"
	"GoSamples/ServiceSample/pkg/config/validation"
)

func ValidateConfig(cfg *types.ServiceConfig) (validation.Warning, error) {
	var (
		warnings validation.Warning
		err      error
	)

	return warnings, err
}
