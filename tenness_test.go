package tennessgo

import "testing"

func checkEquation(got, want interface{}, t *testing.T) {
	if got != want {
		t.Errorf("want %g got %g", want, got)
	}
}

func TestTranslate(t *testing.T) {
	t.Run("test translate word", func(t *testing.T) {
		tr := NewTranslation("bilibili")
		result, err := tr.Translate()
		checkEquation(result, "bilibili", t)
		checkEquation(err, nil, t)
	})
	t.Run("test translate empty", func(t *testing.T) {
		tr := NewTranslation("")
		result, err := tr.Translate()
		checkEquation(result, "", t)
		checkEquation(err.Error(), "empty string", t)
	})
}
