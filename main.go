package main

import "fmt"

func main() {
	// test
	fmt.Println(Reverse("abcd"))
}

// 字符串反转 更改
func Reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
