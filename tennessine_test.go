package tennessine

import "testing"

func checkEquation(got, want interface{}, t *testing.T) {
	if got != want {
		t.Errorf("want %g got %g", want, got)
	}
}

func TestTranslateSingle(t *testing.T) {
	t.Run("test one-to-one", func(*testing.T) {
		word := "没得"
		got, err := TranslateSingle(word)
		want := "没有"
		checkEquation(got, want, t)
		checkEquation(err, nil, t)
	})
	t.Run("test one-to-many", func(t *testing.T) {
		words := []string{"阿婆主", "up主", "Up主"}
		want := "UP主"
		for _, word := range words {
			got, err := TranslateSingle(word)
			checkEquation(got, want, t)
			checkEquation(err, nil, t)
		}
	})
	t.Run("test non-exists", func(t *testing.T) {
		word := "不存在"
		got, err := TranslateSingle(word)
		checkEquation(got, "不存在", t)
		checkEquation(err.Error(), "不存在 is not in the keywords", t)
	})
}
