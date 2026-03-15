package scorer

// ExactScorer performs literal string comparisons.
type ExactScorer struct{}

// Score returns 1.0 if the output matches the expected string exactly, otherwise 0.0.
func (ExactScorer) Score(output, expected string, _ map[string]interface{}) (float64, error) {
	if output == expected {
		return 1.0, nil
	}
	return 0.0, nil
}
