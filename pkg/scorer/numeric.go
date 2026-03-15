package scorer

import (
	"fmt"
	"strconv"
)

// NumericScorer uses a tolerance read from options (default 0.01) to compare numeric outputs.
type NumericScorer struct{}

// Score compares output and expected values within a tolerance.
func (NumericScorer) Score(output, expected string, options map[string]interface{}) (float64, error) {
	tol := 0.01
	if t, ok := options["tolerance"]; ok {
		if v, ok := t.(float64); ok {
			tol = v
		} else {
			return 0.0, fmt.Errorf("invalid tolerance type: %T", t)
		}
	}
	outVal, err := strconv.ParseFloat(output, 64)
	if err != nil {
		return 0.0, err
	}
	exVal, err := strconv.ParseFloat(expected, 64)
	if err != nil {
		return 0.0, err
	}
	if abs(outVal-exVal) <= tol {
		return 1.0, nil
	}
	return 0.0, nil
}

func abs(v float64) float64 {
	if v < 0 {
		return -v
	}
	return v
}
