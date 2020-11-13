package verb

import (
	"RiYu/server/grammar/jiaming"
	"RiYu/server/grammar/static"
)

func yiDuanRoot(verb []rune) (ret []rune, err error) {
	if err = static.VerifyYiDuan(verb); err != nil {
		return
	}
	length := len(verb)
	ret = make([]rune, length-1)
	copy(ret, verb[:length-1])
	return
}

func wuDuanWeiRan(verb []rune) (ret []rune, err error) {
	if err = static.VerifyWuDuan(verb); err != nil {
		return
	}
	length := len(verb)
	ret = make([]rune, length)
	copy(ret, verb[:length-1])
	tail := verb[length-1]
	if tail == 'う' {
		ret[length-1] = 'わ'
	} else {
		ret[length-1] = static.WuduanByU[tail][jiaming.A]
	}
	return
}

func saBianShi(verb []rune) (ret []rune, err error) {
	if err = static.VerifySaBian(verb); err != nil {
		return
	}
	length := len(verb)
	ret = make([]rune, length-1)
	copy(ret, verb[:length-2])
	ret[length-1] = 'し'
	return
}
