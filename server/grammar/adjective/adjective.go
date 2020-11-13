package adjective

import (
	"RiYu/server/grammar/static"
	"fmt"
)

// Forms 形容词各型
type Forms struct {
	Yuan      string `json:"原型"`
	ZhongZhi  string `json:"终止"`
	LianTi    string `json:"连体"`
	MingCi    string `json:"名词"`
	Te        string `json:"て"`
	JieDong   string `json:"接动词"`
	Fou       string `json:"否定"`
	BianXiang string `json:"表变向"`
	TuiLiang  string `json:"推量"`
	JiaDing   string `json:"假定"`
	PanDuan   string `json:"判断"`
}

// GetAdjForms 形容词变型
func GetAdjForms(adj, label string) (ret *Forms, err error) {
	switch label {
	case "形":
		if err = static.VerifyXing(adj); err != nil {
			return
		}
		ret = adjXingForm(adj)
	case "形动":
		ret = adjXingDongForm(adj)
	default:
		err = fmt.Errorf("字段『标』应指定为『形』、『形动』之一，而不应是 %s", label)
	}
	return
}

// Hao 将いい变よい
func Hao(adj string) string {
	if adj == "良い" || adj == "好い" || adj == "いい" {
		return "よい"
	}
	return adj
}

// XingLianYong 形容词连用型
func XingLianYong(word string) []rune {
	adj := []rune(Hao(word))
	length := len(adj)
	ret := make([]rune, length)
	copy(ret, adj[:length-1])
	ret[length-1] = 'く'
	return ret
}

func adjXingForm(adj string) *Forms {
	ah := Hao(adj)
	a := []rune(ah)
	root := string(a[:len(a)-1])
	var panduan string
	if ah == "よい" {
		panduan = "よさ"
	} else if adj == "ない" || adj == "無い" {
		panduan = "なさ"
	} else {
		panduan = root
	}
	return &Forms{
		Yuan:      adj,
		ZhongZhi:  adj,
		LianTi:    adj,
		MingCi:    root + "さ・" + root + "み・" + root + "め",
		Te:        root + "くて",
		JieDong:   root + "く",
		Fou:       root + "くない",
		BianXiang: root + "くなる",
		TuiLiang:  root + "かろう・" + adj + "だろう・" + adj + "でしょう",
		JiaDing:   root + "ければ・" + root + "かったら",
		PanDuan:   panduan + "そう・" + adj + "よう・" + adj + "らしい",
	}
}

func adjXingDongForm(adj string) *Forms {
	return &Forms{
		Yuan:      adj,
		ZhongZhi:  adj + "だ",
		LianTi:    adj + "な",
		MingCi:    adj,
		Te:        adj + "で",
		JieDong:   adj + "に",
		Fou:       adj + "ではない・" + adj + "じゃない",
		BianXiang: adj + "になる",
		TuiLiang:  adj + "だろう・" + adj + "でしょう",
		JiaDing:   adj + "ならば・" + adj + "だったら",
		PanDuan:   adj + "そう・" + adj + "なよう・" + adj + "らしい",
	}
}
