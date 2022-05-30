package main

import (
	// "fmt"
	"testing"
	"unicode/utf8"
)

func TestMain(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"hello,world", "dlrow,olleh"},
		{" ", " "},
		{"!123456", "654321!"},
		// {"张书豪", "豪书张"},
	}
	for _, tc := range testcases {
		rev,_:=Reverse(tc.in)
		if rev!=tc.want {
			t.Errorf("Reverse: %q, want %q", rev, tc.want)
		}
	}
	//fmt.Println("测试成功！")
}
func FuzzMain(f *testing.F) {
	testcases:=[]string{"Hello,World!", " ", "!123456789", "张书豪", "∑σΣ"}
	// testcases:=[]string{"Hello,World!", " ", "!123456789"}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, a string){
		rev,_:=Reverse(a)
		doubleRev,_:=Reverse(rev)
		if a!=doubleRev {
			t.Errorf("Reverse: %q, want %q", a, doubleRev)
		}
		if utf8.ValidString(a)&&!utf8.ValidString(rev) {
			t.Errorf("Reverse 之后是 非utf-8 字符串！%q",rev)
		}
	})
}