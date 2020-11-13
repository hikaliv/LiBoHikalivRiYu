package verb

import "fmt"

// WeiRanFou 动词未然否
func WeiRanFou(verb []rune, label string) (ret []rune, err error) {
	var wr []rune
	switch label {
	case "五段":
		wr, err = wuDuanWeiRan(verb)
	case "一段":
		wr, err = yiDuanRoot(verb)
	case "サ变":
		wr, err = saBianShi(verb)
	case "カ变":
		wr = []rune{'こ'}
	default:
		err = fmt.Errorf("字段『标』应指定为『五段』、『一段』、『サ变』、『カ变』之一，而不应是 %s", label)
	}
	if err == nil {
		ret = append(wr, []rune("ない")...)
	}
	return
}
