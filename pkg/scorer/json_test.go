package scorer

import "testing"

func TestJSONScorer(t *testing.T) {
	s := JSONScorer{}
	if score, err := s.Score(`{"a":1,"b":2}`, `{"b":2,"a":1}`, nil); err != nil || score != 1.0 {
		t.Fatalf("expected JSON match got %v err %v", score, err)
	}
	if score, err := s.Score(`{"a":1}`, `{"a":2}`, nil); err != nil || score != 0.0 {
		t.Fatalf("expected JSON mismatch got %v err %v", score, err)
	}
}
