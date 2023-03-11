package src

import (
	"strings"
)

// Remove chat headers
func RemoveName(ms string) string {
	return strings.ReplaceAll(ms, "@OpenAI", "")
}
