package verb_test

import (
	"RiYu/server/grammar/verb"
	"testing"
)

func TestIKu(test *testing.T) {
	forms, err := verb.GetVerbForms("行く", "五段")
	if err != nil {
		test.Fatal(err)
	}
	ref := verb.Forms{
		Yuan:          "行く",
		LianYong:      "行き",
		LianYongTe:    "行って",
		WeiRanFou:     "行かない",
		WeiRanShiYi:   "行かせる",
		WeiRanBeiDong: "行かれる",
		WeiRanBeiBo:   "行かされる",
		WeiRanKeNeng:  "行ける",
		TuiLiang:      "行こう",
		MingLing:      "行け",
		JiaDing:       "行けば",
		JinZhi:        "行くな",
	}
	if ref != *forms {
		test.Fatalf("『行く』五段变型错误\n应为：\n%v\n却为：\n%v\n", ref, *forms)
	}
}

func TestHaJiMeRu(test *testing.T) {
	forms, err := verb.GetVerbForms("始める", "一段")
	if err != nil {
		test.Fatal(err)
	}
	ref := verb.Forms{
		Yuan:          "始める",
		LianYong:      "始め",
		LianYongTe:    "始めて",
		WeiRanFou:     "始めない",
		WeiRanShiYi:   "始めさせる",
		WeiRanBeiDong: "始められる",
		WeiRanBeiBo:   "始めさせられる",
		WeiRanKeNeng:  "始められる",
		TuiLiang:      "始めよう",
		MingLing:      "始めろ・始めよ",
		JiaDing:       "始めれば",
		JinZhi:        "始めるな",
	}
	if ref != *forms {
		test.Fatalf("『始める』一段变型错误\n应为：\n%v\n却为：\n%v\n", ref, *forms)
	}
}

func TestSuRu(test *testing.T) {
	forms, err := verb.GetVerbForms("する", "サ变")
	if err != nil {
		test.Fatal(err)
	}
	ref := verb.Forms{
		Yuan:          "する",
		LianYong:      "し",
		LianYongTe:    "して",
		WeiRanFou:     "しない・せぬ",
		WeiRanShiYi:   "させる",
		WeiRanBeiDong: "される",
		WeiRanBeiBo:   "させられる",
		WeiRanKeNeng:  "できる",
		TuiLiang:      "しよう",
		MingLing:      "しろ・せよ",
		JiaDing:       "すれば",
		JinZhi:        "するな",
	}
	if ref != *forms {
		test.Fatalf("『する』サ变型错误\n应为：\n%v\n却为：\n%v\n", ref, *forms)
	}
}
