package goRand

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	fmt.Println("生成纯数字随机字符串:")
	fmt.Println(GRands(10, KcRandKindNum))
	fmt.Println("生成小写字母随机字符串:")
	fmt.Println(GRands(10, KcRandKindLower))
	fmt.Println("生成大写字母随机字符串:")
	fmt.Println(GRands(10, KcRandKindUpper))
	fmt.Println("生成特殊字符随机字符串:")
	fmt.Println(GRands(10, KcRandKindSpec))
	fmt.Println("生成所有类型随机字符串:")
	fmt.Println(GRands(10, KcRandKindAll))
}
