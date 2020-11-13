package tense_test

import (
	"RiYu/server/grammar/tense"
	"testing"
)

func TestIKu(test *testing.T) {
	forms, err := tense.GetTenseForms("行く", "五段")
	if err != nil {
		test.Fatal(err)
	}
	ref := tense.Forms{
		Now:        "行く",
		NowJing:    "行きます",
		NoNow:      "行かない",
		NoNowJing:  "行きません",
		Past:       "行った",
		PastJing:   "行きました",
		NoPast:     "行かなかった",
		NoPastJing: "行きませんでした",
	}
	if ref != *forms {
		test.Fatalf("『行く』时态变型错误\n应为：\n%v\n却为：\n%v\n", ref, *forms)
	}
}

func TestII(test *testing.T) {
	forms, err := tense.GetTenseForms("良い", "形")
	if err != nil {
		test.Fatal(err)
	}
	ref := tense.Forms{
		Now:        "良い",
		NowJing:    "良いです",
		NoNow:      "よくない",
		NoNowJing:  "よくないです・よくありません",
		Past:       "よかった",
		PastJing:   "よかったです",
		NoPast:     "よくなかった",
		NoPastJing: "よくなかったです・よくありませんでした",
	}
	if ref != *forms {
		test.Fatalf("『良い』时态变型错误\n应为：\n%v\n却为：\n%v\n", ref, *forms)
	}
}
func TestTaKaI(test *testing.T) {
	forms, err := tense.GetTenseForms("高い", "形")
	if err != nil {
		test.Fatal(err)
	}
	ref := tense.Forms{
		Now:        "高い",
		NowJing:    "高いです",
		NoNow:      "高くない",
		NoNowJing:  "高くないです・高くありません",
		Past:       "高かった",
		PastJing:   "高かったです",
		NoPast:     "高くなかった",
		NoPastJing: "高くなかったです・高くありませんでした",
	}
	if ref != *forms {
		test.Fatalf("『高い』时态变型错误\n应为：\n%v\n却为：\n%v\n", ref, *forms)
	}
}

func TestKiReI(test *testing.T) {
	forms, err := tense.GetTenseForms("綺麗", "形动")
	if err != nil {
		test.Fatal(err)
	}
	ref := tense.Forms{
		Now:        "綺麗だ",
		NowJing:    "綺麗です",
		NoNow:      "綺麗ではない・綺麗じゃない",
		NoNowJing:  "綺麗ではありません・綺麗じゃありません",
		Past:       "綺麗だった",
		PastJing:   "綺麗でした",
		NoPast:     "綺麗ではなかった・綺麗じゃなかった",
		NoPastJing: "綺麗ではありませんでした・綺麗じゃありませんでした",
	}
	if ref != *forms {
		test.Fatalf("『綺麗』时态变型错误\n应为：\n%v\n却为：\n%v\n", ref, *forms)
	}
}
