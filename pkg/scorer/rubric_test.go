package scorer

import "testing"

func TestRubricScorer(t *testing.T) {
	s := RubricScorer{}
	rubric := "foo\nbar"
	if score, err := s.Score("foo baz bar", rubric, nil); err != nil || score <= 0.0 {
		t.Fatalf("expected some score got %v err %v", score, err)
	}
	if score, err := s.Score("none", rubric, nil); err != nil || score != 0.0 {
		t.Fatalf("expected zero score got %v err %v", score, err)
	}
}
