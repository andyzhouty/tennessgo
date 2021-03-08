package tennessine

import (
	"fmt"
	"sort"
)

// all keywords in tennessine
// format {standard: [non-standard-words]}
var keywords = map[string][]string{
	"怎么了": {"肿么了"},
	"没有":  {"没得"},
	"才怪":  {"我信了你的邪"},

	// Keywords from https://bilibili.com/
	"蘑菇弹":     {"蘑菇蛋"},
	"爷的青春回来了": {"爷青回"},
	"爷的青春结束了": {"爷青结"},
	"UP主":     {"up主", "Up主", "阿婆主"},
	"更新视频":    {"鸽视频", "更视频"},
	"字幕君":     {"字母君", "字幕君", "字幕菌", "野生字幕菌"},
	"哈哈哈哈":    {"红红火火恍恍惚惚", "HaHaHa", "HAHAHA"},
	"啊我死了":    {"awsl", "AWSL", "Awsl", "阿伟死了"},
	"笑死我了":    {"xswl", "XSWL", "Xswl", "lol", "LOL", "Lol", "rofl", "Rofl", "ROFL", "Hahaha", "hahaha"},
	"哔哩哔哩":    {"bilibili", "Bilibili", "BiliBili", "BILIBILI"},
	"发生甚么事了":  {"发生肾么事了"},
	"原来是昨天":   {"源赖氏佐田"},
}

func in(target string, array []string) bool {
	sort.Strings(array)
	index := sort.SearchStrings(array, target)
	return index < len(array) && array[index] == target
}

// TranslateSingle translates a single word provided
// it will return the standard mandarin word according to
// the keywords map and a nil value for error
// or it'll return the word itself and a error message like "not-existing is not in the keywords"
func TranslateSingle(word string) (string, error) {
	for k, v := range keywords {
		if in(word, v) {
			return k, nil
		}
	}
	return word, fmt.Errorf("%s is not in the keywords", word)
}
