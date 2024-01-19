package StringUtils

import (
	"regexp"
	"strings"
)

func GenerateSlug(text string) string {
	slug := regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(text, "-")

	slug = strings.ToLower(slug)

	slug = strings.Trim(slug, "-")

	return slug
}
