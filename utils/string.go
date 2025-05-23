package utils

import (
	"strings"
	"unicode"

	"github.com/samber/lo"
)

func NormalizeSpaces(s string) string {
	var builder strings.Builder
	for _, r := range s {
		if unicode.IsSpace(r) {
			builder.WriteRune(' ') // Replace any kind of space with a regular space
		} else {
			builder.WriteRune(r)
		}
	}
	// Collapse multiple spaces and trim
	return strings.Join(strings.Fields(builder.String()), " ")
}

func TemplateReplace(template string, replacements map[string]string) string {
	for key, value := range replacements {
		template = strings.ReplaceAll(template, key, value)
	}
	return template
}

func BetweenLines(content, startLine, endLine string) (string, int, int) {
	lines := strings.Split(content, "\n")
	startIdx := -1
	endIdx := -1
	_, startIdx, _ = lo.FindIndexOf(lines, func(line string) bool {
		return strings.Contains(line, startLine)
	})
	startIdx++
	_, endIdx, _ = lo.FindIndexOf(lines, func(line string) bool {
		return strings.Contains(line, endLine)
	})
	if startIdx < 0 || endIdx < 0 || startIdx > endIdx {
		return "", -1, -1
	}
	return strings.Join(lines[startIdx:endIdx], "\n"), startIdx, endIdx
}
