package verb

import (
	"RiYu/server/grammar/jiaming"
	"RiYu/server/grammar/static"
	"fmt"
)

// Forms 动词各型
type Forms struct {
	Yuan          string `json:"原型"`
	LianYong      string `json:"连用"`
	LianYongTe    string `json:"てた"`
	WeiRanFou     string `json:"未然否定"`
	WeiRanShiYi   string `json:"未然使役"`
	WeiRanBeiDong string `json:"未然被动"`
	WeiRanBeiBo   string `json:"未然被迫"`
	WeiRanKeNeng  string `json:"未然可能"`
	TuiLiang      string `json:"推量"`
	MingLing      string `json:"命令"`
	JiaDing       string `json:"假定"`
	JinZhi        string `json:"禁止"`
}

// GetVerbForms 动词变型
func GetVerbForms(verb, label string) (ret *Forms, err error) {
	switch label {
	case "五段":
		ret, err = verbWuDuanForm(verb)
	case "一段":
		ret, err = verbYiDuanForm(verb)
	case "サ变":
		ret, err = verbSaBianForm(verb)
	case "カ变":
		ret, err = verbKaBianForm(verb)
	default:
		err = fmt.Errorf("字段『标』应指定为『五段』、『一段』、『サ变』、『カ变』之一，而不应是 %s", label)
	}
	return
}

// 似乎五段动词只有 う、つ、る、ぬ、ぶ、む、く、ぐ、す 几种
func verbWuDuanForm(verb string) (forms *Forms, err error) {
	v := []rune(verb)
	if err = static.VerifyWuDuan(v); err != nil {
		return
	}
	lianYong, err := wuDuanLianYong(v)
	if err != nil {
		return
	}
	lianYongTe, voicing, err := wuDuanLianYongTe(v)
	if err != nil {
		return
	}
	if voicing {
		lianYongTe = append(lianYongTe, 'で')
	} else {
		lianYongTe = append(lianYongTe, 'て')
	}
	tail := v[len(v)-1]
	line := static.WuduanByU[tail]
	root := string(v[:len(v)-1])
	wr, err := wuDuanWeiRan(v)
	if err != nil {
		return
	}
	weiran := string(wr)
	var beipo string
	if tail == 'す' {
		beipo = weiran + "せられる"
	} else {
		beipo = weiran + "される"
	}
	forms = &Forms{
		Yuan:          verb,
		LianYong:      string(lianYong),
		LianYongTe:    string(lianYongTe),
		WeiRanFou:     weiran + "ない",
		WeiRanShiYi:   weiran + "せる",
		WeiRanBeiDong: weiran + "れる",
		WeiRanBeiBo:   beipo,
		WeiRanKeNeng:  root + string(line[jiaming.E]) + "る",
		TuiLiang:      root + string(line[jiaming.O]) + "う", // です→でしょう
		MingLing:      root + string(line[jiaming.E]),
		JiaDing:       root + string(line[jiaming.E]) + "ば",
		JinZhi:        verb + "な",
	}
	return
}

func verbYiDuanForm(verb string) (forms *Forms, err error) {
	v := []rune(verb)
	if err = static.VerifyYiDuan(v); err != nil {
		return
	}
	root := string(v[:len(v)-1])
	forms = &Forms{
		Yuan:          verb,
		LianYong:      root,
		LianYongTe:    root + "て",
		WeiRanFou:     root + "ない",
		WeiRanShiYi:   root + "させる",
		WeiRanBeiDong: root + "られる",
		WeiRanBeiBo:   root + "させられる",
		WeiRanKeNeng:  root + "られる",
		TuiLiang:      root + "よう",
		MingLing:      root + "ろ・" + root + "よ",
		JiaDing:       root + "れば",
		JinZhi:        verb + "な",
	}
	return
}

func verbSaBianForm(verb string) (forms *Forms, err error) {
	v := []rune(verb)
	if err = static.VerifySaBian(v); err != nil {
		return
	}
	root := string(v[:len(v)-2])
	forms = &Forms{
		Yuan:          verb,
		LianYong:      root + "し",
		LianYongTe:    root + "して",
		WeiRanFou:     root + "しない・" + root + "せぬ",
		WeiRanShiYi:   root + "させる",
		WeiRanBeiDong: root + "される",
		WeiRanBeiBo:   root + "させられる",
		WeiRanKeNeng:  root + "できる",
		TuiLiang:      root + "しよう",
		MingLing:      root + "しろ・" + root + "せよ",
		JiaDing:       root + "すれば",
		JinZhi:        verb + "な",
	}
	return
}

func verbKaBianForm(verb string) (forms *Forms, err error) {
	if err = static.VerifyKaBian(verb); err != nil {
		return
	}
	forms = &Forms{
		Yuan:          verb,
		LianYong:      "き",
		LianYongTe:    "きて",
		WeiRanFou:     "こない",
		WeiRanShiYi:   "こさせる",
		WeiRanBeiDong: "こられる",
		WeiRanBeiBo:   "こさせられる",
		WeiRanKeNeng:  "こられる",
		TuiLiang:      "こよう",
		MingLing:      "こい",
		JiaDing:       "くれば",
		JinZhi:        verb + "な",
	}
	return
}
