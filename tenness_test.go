package tennessgo

import "testing"

func checkEquation(got, want interface{}, t *testing.T) {
	if got != want {
		t.Errorf("want %v got %v", want, got)
	}
}

func TestTranslateWordOrEmpty(t *testing.T) {
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
		checkEquation(err.Error(), "empty string to translate", t)
	})
}

func TestTranslateSentence(t *testing.T) {
	t.Run("without reserved keywords", func(t *testing.T) {
		tr := NewTranslation("咋给轮胎冲气")
		result, err := tr.Translate()
		checkEquation(result, "怎么给轮胎充气", t)
		checkEquation(err, nil, t)
	})
	t.Run("with reserved keywords", func(t *testing.T) {
		tr := NewTranslation("发生甚么事了是啥意思")
		result, err := tr.Translate()
		checkEquation(result, "发生甚么事了是什么意思", t)
		checkEquation(err, nil, t)
	})
	t.Run("avoid translating specific format", func(t *testing.T) {
		tr := NewTranslation("{k@#219}发生甚么事了是啥意思")
		result, err := tr.Translate()
		checkEquation(result, "{k@#219}发生甚么事了是什么意思", t)
		checkEquation(err, nil, t)
	})
	t.Run("test one keyword is IN another one", func(t *testing.T) {
		tr := NewTranslation("马可·波萝")
		result, _ := tr.Translate()
		checkEquation(result, "马可·波罗", t)
	})
}
