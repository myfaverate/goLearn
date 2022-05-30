package main

import "fmt"

func main() {
	// test
	input:="abcdef省赛ghigklmn"
	rev:=Reverse(input)
	doubleRev:=Reverse(rev)
	fmt.Printf("原始字符串: %q\n", input)
	fmt.Printf("反转之后的字符串%q\n",rev)
	fmt.Printf("双重反转之后的字符串%q\n",doubleRev)
}

// 字符串反转 更改
// 第三次更改
func Reverse(s string) string {
	b := []rune(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
