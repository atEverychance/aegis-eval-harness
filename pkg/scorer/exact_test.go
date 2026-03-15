package scorer

import "testing"

func TestExactScorer(t *testing.T) {
	s := ExactScorer{}
	if score, err := s.Score("hello", "hello", nil); err != nil || score != 1.0 {
		t.Fatalf("expected 1.0 got %v err %v", score, err)
	}
	if score, err := s.Score("no", "yes", nil); err != nil || score != 0.0 {
		t.Fatalf("expected 0.0 got %v err %v", score, err)
	}
}
