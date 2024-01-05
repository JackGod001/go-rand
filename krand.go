package goRand

import (
	"math/rand"
	"time"
)

const (
	KcRandKindNum   = 0 // 纯数字
	KcRandKindLower = 1 // 小写字母
	KcRandKindUpper = 2 // 大写字母
	KcRandKindSpec  = 3 // 特殊字符
	KcRandKindAll   = 4 // 数字、大小写字母、特殊字符
)

// 特殊字符集
var specialChars = []int{33, 35, 36, 37, 38, 40, 41, 42, 43, 45, 46, 47, 58, 59, 60, 61, 62, 63, 64, 91, 92, 93, 94, 95, 96, 123, 124, 125, 126}

// GRands 随机字符串
func GRands(size int, kind int) string {
	ikind, kinds, result := kind, [][]int{
		{10, 48},                // 数字
		{26, 97},                // 小写字母
		{26, 65},                // 大写字母
		{len(specialChars), -1}, // 特殊字符，-1是特殊标记
	}, make([]byte, size)
	isAll := kind == KcRandKindAll
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if isAll { // random ikind
			ikind = rand.Intn(4) // 现在有4种类型
		}
		if ikind == KcRandKindSpec { // 特殊字符处理
			result[i] = uint8(specialChars[rand.Intn(len(specialChars))])
		} else {
			scope, base := kinds[ikind][0], kinds[ikind][1]
			result[i] = uint8(base + rand.Intn(scope))
		}
	}
	return string(result)
}
