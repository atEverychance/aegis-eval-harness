package scorer

import "strings"

// RubricScorer matches key phrases from a rubric.
type RubricScorer struct{}

// Score evaluates how many phrases the output satisfies versus the rubric.
func (RubricScorer) Score(output, expected string, options map[string]interface{}) (float64, error) {
	phrases := extractRubric(expected)
	if len(phrases) == 0 {
		return 0.0, nil
	}
	matches := 0
	for _, phrase := range phrases {
		if strings.Contains(output, phrase) {
			matches++
		}
	}
	return float64(matches) / float64(len(phrases)), nil
}

func extractRubric(raw string) []string {
	lines := strings.Split(raw, "\n")
	result := make([]string, 0, len(lines))
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}
