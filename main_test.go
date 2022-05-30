package main

import (
	// "fmt"
	"testing"
)

func TestMain(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"hello,world", "dlrow,olleh"},
		{" ", " "},
		{"!123456", "654321!"},
		{"张书豪", "豪书张"},
	}
	for _, tc := range testcases {
		rev:=Reverse(tc.in)
		if rev!=tc.want {
			t.Errorf("Reverse: %q, want %q", rev, tc.want)
		}
	}
	//fmt.Println("测试成功！")
}