package verb

import (
	"RiYu/server/grammar/jiaming"
	"RiYu/server/grammar/static"
	"fmt"
)

// LianYong 动词连用型
func LianYong(verb []rune, label string) (ret []rune, err error) {
	switch label {
	case "五段":
		ret, err = wuDuanLianYong(verb)
	case "一段":
		ret, err = yiDuanRoot(verb)
	case "サ变":
		ret, err = saBianShi(verb)
	case "カ变":
		ret = []rune{'き'}
	default:
		err = fmt.Errorf("字段『标』应指定为『五段』、『一段』、『サ变』、『カ变』之一，而不应是 %s", label)
	}
	return
}

// LianYongTe 动词连用型
func LianYongTe(verb []rune, label string) (ret []rune, voicing bool, err error) {
	switch label {
	case "五段":
		ret, voicing, err = wuDuanLianYongTe(verb)
	case "一段", "サ变", "カ变":
		ret, err = LianYong(verb, label)
	default:
		err = fmt.Errorf("字段『标』应指定为『五段』、『一段』、『サ变』、『カ变』之一，而不应是 %s", label)
	}
	return
}

func wuDuanLianYong(verb []rune) (ret []rune, err error) {
	if err = static.VerifyWuDuan(verb); err != nil {
		return
	}
	length := len(verb)
	line := static.WuduanByU[verb[length-1]]
	ret = make([]rune, length)
	copy(ret, verb[:length-1])
	ret[length-1] = line[jiaming.I]
	return
}

func wuDuanLianYongTe(verb []rune) (ret []rune, voicing bool, err error) {
	if err = static.VerifyWuDuan(verb); err != nil {
		return
	}
	length := len(verb)
	ret = make([]rune, length)
	copy(ret, verb[:length-1])
	if (verb[length-2] == '行' || verb[length-2] == 'い') && verb[length-1] == 'く' {
		ret[length-1] = 'っ'
	} else {
		switch verb[length-1] {
		case 'う', 'つ', 'る':
			ret[length-1] = 'っ'
		case 'ぬ', 'ぶ', 'む':
			ret[length-1] = 'ん'
			voicing = true // 浊化
		case 'く':
			ret[length-1] = 'い'
		case 'ぐ':
			ret[length-1] = 'い'
			voicing = true
		case 'す':
			ret[length-1] = 'し'
		default:
			err = fmt.Errorf("无法以五段动词变处理 %s", string(verb))
		}
	}
	return
}
