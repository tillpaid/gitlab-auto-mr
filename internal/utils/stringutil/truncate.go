package stringutil

import (
	"strings"
)

func TruncateWords(s string, max int) string {
	if len(s) <= max {
		return s
	}

	words := strings.Fields(s)
	var result []string
	length := 0

	for i, word := range words {
		wordLen := len(word)
		if i > 0 {
			wordLen++
		}

		if length+wordLen > max {
			break
		}

		result = append(result, word)
		length += wordLen
	}

	return strings.Join(result, " ") + "..."
}
