package resource

import (
	"github.com/WebmiRU/YudolePlatformPackages/theme"
)

type Theme struct {
	Payload *theme.Theme `json:"payload"`
}

type ThemeIndex struct {
	Payload map[string]*theme.Theme `json:"payload"`
}
