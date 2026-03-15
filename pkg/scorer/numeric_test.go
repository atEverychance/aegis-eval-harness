package scorer

import "testing"

func TestNumericScorer(t *testing.T) {
	s := NumericScorer{}
	if score, err := s.Score("1.0", "1.005", map[string]interface{}{"tolerance": 0.01}); err != nil || score != 1.0 {
		t.Fatalf("expected within tolerance got %v err %v", score, err)
	}
	if score, err := s.Score("1.0", "1.02", map[string]interface{}{"tolerance": 0.01}); err != nil || score != 0.0 {
		t.Fatalf("expected outside tolerance got %v err %v", score, err)
	}
}
