package static

import (
	"RiYu/server/grammar/jiaming"
	"fmt"
)

// WuduanByU う段索引之音行
var WuduanByU = jiaming.WuDuanYinTu(jiaming.U)

// VerifyWuDuan 验是否五段动词
func VerifyWuDuan(verb []rune) error {
	_, ok := WuduanByU[verb[len(verb)-1]]
	if !ok {
		return fmt.Errorf("%s 非五段动词", string(verb))
	}
	return nil
}

// VerifyYiDuan 验是否一段动词
func VerifyYiDuan(verb []rune) error {
	if verb[len(verb)-1] != 'る' {
		return fmt.Errorf("%s 非一段动词", string(verb))
	}
	return nil
}

// VerifySaBian 验是否サ变动词
func VerifySaBian(verb []rune) error {
	if verb[len(verb)-2] != 'す' || verb[len(verb)-1] != 'る' {
		return fmt.Errorf("%s 非サ变动词", string(verb))
	}
	return nil
}

// VerifyKaBian 验是否カ变动词
func VerifyKaBian(verb string) error {
	if verb != "来る" && verb != "くる" {
		return fmt.Errorf("%s 非カ变动词", verb)
	}
	return nil
}

// VerifyXing 验是否形容词
func VerifyXing(word string) error {
	adj := []rune(word)
	if adj[len(adj)-1] != 'い' {
		return fmt.Errorf("%s 非形容词", word)
	}
	return nil
}
