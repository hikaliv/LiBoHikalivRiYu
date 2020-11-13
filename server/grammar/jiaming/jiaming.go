package jiaming

// JiaMing 假名
var JiaMing = []rune("あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもらりるれろがぎぐげござじずぜぞだぢづでどばびぶべぼぱぴぷぺぽやゆよわをんっゃゅょ")

// WuDuanBiao 五段标
type WuDuanBiao int

// 五段标值
const (
	A WuDuanBiao = iota
	I
	U
	E
	O
)

// WuDuanYinTu 五段音图典，只包含所有满五段的行
func WuDuanYinTu(key WuDuanBiao) map[rune]map[WuDuanBiao]rune {
	// 有效一共 13 行 65 个假名
	jiaming := JiaMing[:len(JiaMing)-10]
	length := len(jiaming)
	yintu := make(map[rune]map[WuDuanBiao]rune, length/5)
	for itor := 0; itor < length; itor += 5 {
		r := jiaming[itor+int(key)]
		yintu[r] = make(map[WuDuanBiao]rune, 5)
		line := yintu[r]
		line[A] = jiaming[itor]
		line[I] = jiaming[itor+1]
		line[U] = jiaming[itor+2]
		line[E] = jiaming[itor+3]
		line[O] = jiaming[itor+4]
	}
	return yintu
}
