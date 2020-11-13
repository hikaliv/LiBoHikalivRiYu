package adjective_test

import (
	"RiYu/server/grammar/adjective"
	"testing"
)

func TestII(test *testing.T) {
	forms, err := adjective.GetAdjForms("いい", "形")
	if err != nil {
		test.Fatal(err)
	}
	ref := adjective.Forms{
		Yuan:      "いい",
		ZhongZhi:  "いい",
		LianTi:    "いい",
		MingCi:    "よさ・よみ・よめ",
		Te:        "よくて",
		JieDong:   "よく",
		Fou:       "よくない",
		BianXiang: "よくなる",
		TuiLiang:  "よかろう・いいだろう・いいでしょう",
		JiaDing:   "よければ・よかったら",
		PanDuan:   "よさそう・いいよう・いいらしい",
	}
	if ref != *forms {
		test.Fatalf("『いい』变型错误\n应为：\n%v\n却为：\n%v\n", ref, *forms)
	}
}

func TestSuKi(test *testing.T) {
	forms, err := adjective.GetAdjForms("好き", "形动")
	if err != nil {
		test.Fatal(err)
	}
	ref := adjective.Forms{
		Yuan:      "好き",
		ZhongZhi:  "好きだ",
		LianTi:    "好きな",
		MingCi:    "好き",
		Te:        "好きで",
		JieDong:   "好きに",
		Fou:       "好きではない・好きじゃない",
		BianXiang: "好きになる",
		TuiLiang:  "好きだろう・好きでしょう",
		JiaDing:   "好きならば・好きだったら",
		PanDuan:   "好きそう・好きなよう・好きらしい",
	}
	if ref != *forms {
		test.Fatalf("『好き』变型错误\n应为：\n%v\n却为：\n%v\n", ref, *forms)
	}
}
