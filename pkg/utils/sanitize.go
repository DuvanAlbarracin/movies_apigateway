package utils

import (
	"html/template"
	"strings"
)

func SanitizeString(s string) string {
	return template.HTMLEscapeString(s)
}

func TrimString(s string) string {
	return strings.TrimRight(s, "\r\n")
}
