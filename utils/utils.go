package utils

import (
	"github.com/google/uuid"
	"regexp"
	"strings"
)

func RemoveNewlines(input string) string {
	return strings.ReplaceAll(input, "\n", "")
}

func GenerateUniqueID() string {
	return uuid.New().String()
}

func SplitResponse(response string) []string {
	re := regexp.MustCompile(`\d+\.`)
	points := re.Split(response, -1)

	var result []string
	for _, point := range points {
		if point != "" {
			result = append(result, strings.TrimSpace(point))
		}
	}

	return result
}
