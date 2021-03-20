package tennessgo

import (
	"fmt"
	"testing"
)

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

func TestTranslateDeclarativeSentence(t *testing.T) {
	t.Run("without reserved keywords", func(t *testing.T) {
		tr := NewTranslation("给轮胎冲气")
		result, err := tr.Translate()
		checkEquation(result, "给轮胎充气", t)
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
}

func TestWhatQuestions(t *testing.T) {
	t.Run("questions with question marks", func(t *testing.T) {
		tr := NewTranslation("仃车是什么意思？")
		tr2 := NewTranslation("仃车是什么意思?")
		result, err := tr.Translate()
		result2, _ := tr2.Translate()
		checkEquation(result, "仃车是什么意思", t)
		checkEquation(result2, "仃车是什么意思", t)
		checkEquation(err, nil, t)
	})
	t.Run("asking what-meaning questions", func(t *testing.T) {
		tr := NewTranslation("合饭是什么东西")
		tr2 := NewTranslation("合饭是啥玩意儿")
		tr3 := NewTranslation("合饭是什么玩意儿")
		result, err := tr.Translate()
		result2, _ := tr2.Translate()
		result3, _ := tr3.Translate()
		checkEquation(err, nil, t)
		checkEquation(result, "合饭是什么", t)
		checkEquation(result2, "合饭是什么", t)
		checkEquation(result3, "合饭是什么", t)
	})
}

func ExampleTranslate() {
	tr := Translate{ToTranslate: "发生甚么事了是啥意思", ReservedKeywords: ReservedKeywords}
	fmt.Println(tr.ToTranslate)
	// output:
	// 发生甚么事了是啥意思
}

func ExampleNewTranslation() {
	tr := NewTranslation("")
	fmt.Println(tr.ReservedKeywords[227])
	// output:
	// 耗子尾汁
}

func ExampleTranslate_Translate() {
	tr := NewTranslation("发生甚么事了是啥意思")
	result, err := tr.Translate()
	fmt.Println(result, err)
	// output:
	// 发生甚么事了是什么意思 <nil>
}
