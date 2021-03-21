package tennessgo

import (
	"fmt"
	"testing"
)

func assertEqual(got, want interface{}, t *testing.T) {
	if got != want {
		t.Errorf("want %v got %v", want, got)
	}
}

func TestTranslateWordOrEmpty(t *testing.T) {
	t.Run("test translate word", func(t *testing.T) {
		tr := NewTranslation("bilibili")
		result, err := tr.Translate()
		assertEqual(result, "bilibili", t)
		assertEqual(err, nil, t)
	})
	t.Run("test translate empty", func(t *testing.T) {
		tr := NewTranslation("")
		result, err := tr.Translate()
		assertEqual(result, "", t)
		assertEqual(err.Error(), "empty string to translate", t)
	})
}

func TestTranslateDeclarativeSentence(t *testing.T) {
	t.Run("without reserved keywords", func(t *testing.T) {
		tr := NewTranslation("给轮胎冲气")
		result, err := tr.Translate()
		assertEqual(result, "给轮胎充气", t)
		assertEqual(err, nil, t)
	})
	t.Run("with reserved keywords", func(t *testing.T) {
		tr := NewTranslation("发生甚么事了")
		result, err := tr.Translate()
		assertEqual(result, "发生甚么事了", t)
		assertEqual(err, nil, t)
	})
	t.Run("avoid translating specific format", func(t *testing.T) {
		tr := NewTranslation("{k@#219}发生甚么事了是啥意思")
		result, err := tr.Translate()
		assertEqual(result, "{k@#219}发生甚么事了是什么意思", t)
		assertEqual(err, nil, t)
	})
}

func TestWhatQuestions(t *testing.T) {
	t.Run("toTranslate only contains a question mark", func(t *testing.T) {
		tr := NewTranslation("？")
		result, err := tr.Translate()
		assertEqual(result, "", t)
		assertEqual(err.Error(), "translating a string only contains a question mark", t)
	})
	t.Run("questions with question marks", func(t *testing.T) {
		tr := NewTranslation("仃车是什么意思？")
		tr2 := NewTranslation("仃车是什么意思?")
		result, err := tr.Translate()
		result2, _ := tr2.Translate()
		assertEqual(result, "仃车是什么意思?", t)
		assertEqual(result2, "仃车是什么意思?", t)
		assertEqual(err, nil, t)
	})
	t.Run("asking what-meaning questions", func(t *testing.T) {
		tr := NewTranslation("合饭是什么东西")
		tr2 := NewTranslation("合饭是啥玩意儿")
		tr3 := NewTranslation("合饭是什么玩意儿")
		result, err := tr.Translate()
		result2, _ := tr2.Translate()
		result3, _ := tr3.Translate()
		assertEqual(err, nil, t)
		assertEqual(result, "合饭是什么", t)
		assertEqual(result2, "合饭是什么", t)
		assertEqual(result3, "合饭是什么", t)
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
