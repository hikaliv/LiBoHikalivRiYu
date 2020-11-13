package tense

import (
	"RiYu/server/grammar/adjective"
	"RiYu/server/grammar/static"
	"RiYu/server/grammar/verb"
	"fmt"
)

// Forms 时态各型
type Forms struct {
	Now        string `json:"现时"`
	NowJing    string `json:"现敬"`
	NoNow      string `json:"现否"`
	NoNowJing  string `json:"现否敬"`
	Past       string `json:"既时"`
	PastJing   string `json:"既敬"`
	NoPast     string `json:"既否"`
	NoPastJing string `json:"既否敬"`
}

// GetTenseForms 动词、形容词时态变型
func GetTenseForms(word, label string) (ret *Forms, err error) {
	switch label {
	case "五段", "一段", "サ变", "カ变":
		ret, err = tenseDong(word, label)
	case "形":
		if err = static.VerifyXing(word); err != nil {
			return
		}
		ret = tenseXing(word)
	case "形动":
		ret = tenseXingDong(word)
	default:
		err = fmt.Errorf("字段『标』应指定为『五段』、『一段』、『サ变』、『カ变』、『形』、『形动』之一，而不应是 %s", label)
	}
	return
}

func tenseDong(word, label string) (ret *Forms, err error) {
	v := []rune(word)
	masu, err := masu(v, label)
	if err != nil {
		return
	}
	noNow, err := verb.WeiRanFou(v, label)
	if err != nil {
		return
	}
	past, voicing, err := verb.LianYongTe(v, label)
	if voicing {
		past = append(past, 'だ')
	} else {
		past = append(past, 'た')
	}
	noPast := adjNowToPast(noNow)
	ret = &Forms{
		Now:        word,
		NowJing:    string(masu),
		NoNow:      string(noNow),
		NoNowJing:  string(masuToNo(masu)),
		Past:       string(past),
		PastJing:   string(masuToPast(masu)),
		NoPast:     string(noPast),
		NoPastJing: string(masuToNoPast(masu)),
	}
	return
}

func tenseXing(word string) *Forms {
	desu := func(word string) string {
		return word + "です"
	}
	lianyong := adjective.XingLianYong(word)
	noNow := append(lianyong, 'な', 'い')
	noNowStr := string(noNow)
	past := string(adjNowToPast([]rune(adjective.Hao(word))))
	noPast := string(adjNowToPast(noNow))
	noJing := "・" + string(lianyong) + "ありません"
	return &Forms{
		Now:        word,
		NowJing:    desu(word),
		NoNow:      noNowStr,
		NoNowJing:  desu(noNowStr) + noJing,
		Past:       past,
		PastJing:   desu(past),
		NoPast:     noPast,
		NoPastJing: desu(noPast) + noJing + "でした",
	}
}

func tenseXingDong(word string) *Forms {
	return &Forms{
		Now:        word + "だ",
		NowJing:    word + "です",
		NoNow:      word + "ではない・" + word + "じゃない",
		NoNowJing:  word + "ではありません・" + word + "じゃありません",
		Past:       word + "だった",
		PastJing:   word + "でした",
		NoPast:     word + "ではなかった・" + word + "じゃなかった",
		NoPastJing: word + "ではありませんでした・" + word + "じゃありませんでした",
	}
}

func adjNowToPast(adj []rune) []rune {
	length := len(adj)
	past := make([]rune, length+2)
	copy(past, adj[:length-1])
	copy(past[length-1:], []rune("かった"))
	return past
}

func masu(word []rune, label string) (ret []rune, err error) {
	lianyong, err := verb.LianYong(word, label)
	if err != nil {
		return
	}
	ret = append(lianyong, []rune("ます")...)
	return
}

func masuToPast(masu []rune) []rune {
	length := len(masu)
	past := make([]rune, length+1)
	copy(past, masu[:length-2])
	copy(past[length-2:], []rune("ました"))
	return past
}

func masuToNo(masu []rune) []rune {
	length := len(masu)
	no := make([]rune, length+1)
	copy(no, masu[:length-2])
	copy(no[length-2:], []rune("ません"))
	return no
}

func masuToNoPast(masu []rune) []rune {
	length := len(masu)
	pastNo := make([]rune, length+4)
	copy(pastNo, masu[:length-2])
	copy(pastNo[length-2:], []rune("ませんでした"))
	return pastNo
}
