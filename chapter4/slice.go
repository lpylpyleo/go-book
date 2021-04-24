package main

import (
	"fmt"
	"unicode"
)

func main() {
	// a := [...]int{1, 2, 3, 4, 5}
	// reverse(a[:])
	// fmt.Println(a)

	// var x, y []int
	// for i := 0; i < 10; i++ {
	// 	y = appendInt(x, i)
	// 	fmt.Printf("%d  cap=%d\t%v\n", i, cap(y), y)
	// 	x = y
	// }
	// s := []int{5, 6, 7, 8, 9}
	// fmt.Println(remove(s, 2))

	// s := []string{"a", "b", "c", "c", "d", "d"}
	// zs := zipSame(s)
	// fmt.Println(zs)

	s := []byte{' ', 'a', 'd', ' ', ' ', ' ', 'q'}
	zs := zipSpace(s)
	fmt.Printf("%v", zs)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func appendInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}

// 不保持原来顺序删除
func remove(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func reverse2(arr *[10]int) {}

func zipSame(s []string) []string {
	if len(s) <= 1 {
		return s
	}
	r := s[:1]
	for i := 1; i < len(s)-1; i++ {
		if s[i-1] != s[i] {
			r = append(r, s[i])
		}
	}
	return r
}

func zipSpace(s []byte) []byte {
	if len(s) <= 1 {
		return s
	}
	var offset int
	for i := 1; i < len(s)-1; i++ {
		fmt.Println(rune(s[i]), rune(s[i-1]))
		if unicode.IsSpace(rune(s[i])) && unicode.IsSpace(rune(s[i-1])) {
			copy(s[i:], s[i+1:])
			offset++
			i--
		}
	}
	return s[:len(s)-offset]
}
