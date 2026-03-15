package scorer

import "encoding/json"

// JSONScorer performs deep JSON comparisons.
type JSONScorer struct{}

// Score unmarshals both strings and compares their normalized representations.
func (JSONScorer) Score(output, expected string, _ map[string]interface{}) (float64, error) {
	var out, exp interface{}
	if err := json.Unmarshal([]byte(output), &out); err != nil {
		return 0.0, err
	}
	if err := json.Unmarshal([]byte(expected), &exp); err != nil {
		return 0.0, err
	}
	return compareJSON(out, exp), nil
}

func compareJSON(a, b interface{}) float64 {
	switch x := a.(type) {
	case map[string]interface{}:
		if y, ok := b.(map[string]interface{}); ok {
			if len(x) != len(y) {
				return 0.0
			}
			for key, val := range x {
				if res := compareJSON(val, y[key]); res == 0.0 {
					return 0.0
				}
			}
			return 1.0
		}
	case []interface{}:
		if y, ok := b.([]interface{}); ok {
			if len(x) != len(y) {
				return 0.0
			}
			for i := range x {
				if res := compareJSON(x[i], y[i]); res == 0.0 {
					return 0.0
				}
			}
			return 1.0
		}
	default:
		if a == b {
			return 1.0
		}
	}
	return 0.0
}
