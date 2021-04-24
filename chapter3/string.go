package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	// fmt.Println(printInts([]int{2, 4, 5}))
	fmt.Println(comma2("abcdefg"))
}

func basename1(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot > 0 {
		s = s[:dot]
	}
	return s
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return s[:n-3] + "," + s[n-3:]
}

func printInts(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteByte(',')
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func comma2(s string) string {
	var buf bytes.Buffer
	l := len(s)

	for i := 0; i < l; i++ {
		if (l-i)%3 == 0 && i != 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}

	return buf.String()
}

func isMessed(s1, s2 string) bool {
	if len(s1) != len(s2) || s1 == s2 {
		return false
	}

	a1, a2 := strings.Split(s1, ""), strings.Split(s2, "")

	fmt.Println(a1, a2) // 排序遍历一一对比。。。

	return true
}
