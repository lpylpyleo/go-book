package main

import (
	"fmt"
	"strings"
)

func main() {
	// matrix := [][]int{
	// 	{1, 2, 3, 4},
	// 	{5, 6, 7, 8},
	// 	{9, 10, 11, 12},
	// 	{13, 14, 15, 16},
	// }
	// matrix := [][]int{
	// 	{1, 2, 3},
	// 	{5, 6, 7},
	// 	{9, 10, 11},
	// }
	// rotate(matrix)
	// fmt.Println(matrix)
	// fmt.Println(index2cood(2, 2, 1, 1))
	// fmt.Println(cood2index(1, -1, 1, 1))
	// fmt.Println(longestPalindrome("babad"))
	// fmt.Println(longestPalindrome("cbbd"))
	// fmt.Println(longestPalindrome("ac"))
	// fmt.Println(longestPalindrome("abcbdedbaa"))
	// fmt.Println(longestPalindrome("bb"))
	// fmt.Println(longestPalindrome("ccc"))
	// fmt.Println(longestPalindrome("aacabdkacaa"))
	// fmt.Println(jump([]int{5, 6, 4, 4, 6, 9, 4, 4, 7, 4, 4, 8, 2, 6, 8, 1, 5, 9, 6, 5, 2, 7, 9, 7, 9, 6, 9, 4, 1, 6, 8, 8, 4, 4, 2, 0, 3, 8, 5}))
	// fmt.Println(jump([]int{1, 2, 1, 1, 1}))
	// fmt.Println(jump([]int{2, 3, 1, 1, 4}))
	// fmt.Println(jump([]int{0}))
	// var a int = 1
	// a |= 0
	// fmt.Println(a)
	// fmt.Println(isMatch("aa", "*"))
	// fmt.Println(isMatch("cb", "?a"))
	// fmt.Println(isMatch("adceb", "*a*b"))
	// fmt.Println(isMatch("dceb", "*b"))
	// fmt.Println(isMatch("acdcb", "a*c?b"))
	// fmt.Println(isMatch("c", "*?*"))
	// fmt.Println(isMatch("bbbababbbbabbbbababbaaabbaababbbaabbbaaaabbbaaaabb", "*b********bb*b*bbbbb*ba"))
	// fmt.Println(isMatch("abbabaaabbabbaababbabbbbbabbbabbbabaaaaababababbbabababaabbababaabbbbbbaaaabababbbaabbbbaabbbbababababbaabbaababaabbbababababbbbaaabbbbbabaaaabbababbbbaababaabbababbbbbababbbabaaaaaaaabbbbbaabaaababaaaabb", "**aa*****ba*a*bb**aa*ab****a*aaaaaa***a*aaaa**bbabb*b*b**aaaaaaaaa*a********ba*bbb***a*ba*bb*bb**a*b*bb"))

	// fmt.Println(isMatch("abcabczzzde", "*abc???de*"))
	// fmt.Println(removeDupAsteria("*b********bb*b*bbbbb*ba"))

	src := []int{36, 34, 345, 13, 35, 4, 54, 225, 34, 23, 6, 23}
	// src := []int{1, 43, 64, 23, 23, 54, 2, 67, 4, 46, 33}
	quickSort(src)
	fmt.Println(src)
}

func permute(nums []int) [][]int {
	var result [][]int
	permuteHelper(nums, nil, &result)
	return result
}

func permuteHelper(left, route []int, result *[][]int) {
	ll := len(left)
	if ll == 1 {
		route = append(route, left[0])
		*result = append(*result, route)
		return
	}

	for i, v := range left {
		//??????
		newRoute := make([]int, len(route))
		copy(newRoute, route)
		newLeft := make([]int, ll)
		copy(newLeft, left)

		// ??????route?????????????????????
		newRoute = append(newRoute, v)

		//????????????????????????left
		newLeft[i] = newLeft[ll-1]
		permuteHelper(newLeft[:ll-1], newRoute, result)
	}
}

func expand(s string, f func(string) string) string {
	return strings.Replace(s, "foo", f("foo"), -1)
}

// func permute(nums []int) [][]int {
// 	var res [][]int

// 	var add func(n *Node, nums []int) *[]Node
// 	var dump func(seen []int, n *Node)

// 	add = func(n *Node, nums []int) *[]Node {
// 		var children []Node
// 		if n.children == nil {
// 			n.children = &[]Node{}
// 		}
// 		for i, num := range nums {
// 			cp := make([]int, len(nums))
// 			copy(cp, nums)
// 			cp[i] = cp[len(cp)-1]
// 			var c Node
// 			c = Node{num, add(&c, cp[:len(cp)-1])}
// 			children = append(children, c)
// 			*n.children = append(*n.children, c)
// 		}
// 		return &children
// 	}

// 	dump = func(seen []int, n *Node) {
// 		seen = append(seen, n.value)
// 		if len(*n.children) == 0 {
// 			res = append(res, seen[1:])
// 		}

// 		for _, c := range *n.children {
// 			seenCpy := make([]int, len(seen))
// 			copy(seenCpy, seen)
// 			dump(seenCpy, &c)
// 		}
// 	}

// 	root := &Node{0, &[]Node{}}
// 	add(root, nums)
// 	dump(nil, root)
// 	return res
// }

type Node struct {
	value    int
	children *[]Node
}

type Point struct {
	X, Y int
}

func rotate(matrix [][]int) {
	// dimension
	n := len(matrix)

	// init seen
	seen := make([][]bool, n)
	for i := 0; i < n; i++ {
		seen[i] = make([]bool, n)
	}

	fc := float64(n)/2.0 - 0.5

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			x, y := index2cood(i, j, fc, fc)
			cood := complex(float64(x), float64(y))
			rotatedCood := cood * complex(0, -1)
			ri, rj := cood2index(real(rotatedCood), imag(rotatedCood), fc, fc)
			if !seen[i][j] && !seen[ri][rj] {
				matrix[i][j], matrix[ri][rj] = matrix[ri][rj], matrix[i][j]
				seen[ri][rj] = true
			}
		}
	}
}

// ???????????????
func index2cood(i, j int, cx, cy float64) (float64, float64) {
	return float64(j) - cy, cx - float64(i)
}

// ???????????????
func cood2index(x, y, cx, cy float64) (int, int) {
	return int(cx - y), int(x + cy)
}

func longestPalindrome(s string) string {
	b := []byte(s)
	l := len(b)
	if l == 1 {
		return s
	}
	result := b[:0]

	for i := 0; i < l; i++ {
		left := 0
		right := 0
		size := 0
		for j := 0; j <= i && j < l-i; j++ {
			if i+j+1 < l {
				if b[i-j] == b[i+j+1] {
					left = i - j
					right = i + j + 1
				} else {
					break
				}
			}
		}
		for j := 0; j <= i && j < l-i; j++ {
			if b[i-j] == b[i+j] {
				size = j
			} else {
				break
			}
		}
		if right-left+1 > len(result) {
			result = b[left : right+1]
		}
		if size*2+1 > len(result) {
			result = b[i-size : i+size+1]
		}
	}

	return string(result)
}

func isPalindrome(s []byte) bool {
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i] {
			return false
		}
	}
	return true
}

func jump(nums []int) int {
	var steps int
	jumpHelper(nums, &steps)
	return steps
}

func jumpHelper(nums []int, steps *int) {
	l := len(nums)
	// Already at end point
	if l <= 1 {
		return
	}
	// Able to jump to end point
	if nums[0]+1 >= l {
		*steps++
		return
	}

	var next int //????????????????????????
	var max int  //??????????????????????????????
	for i := 0; i < nums[0]; i++ {
		dist := nums[i+1] + i
		if dist > max {
			max = dist
			next = i
		}
	}
	*steps++
	jumpHelper(nums[next+1:], steps)
}

const (
	asteria  = '*'
	question = '?'
)

// TODO ????????????????????????????????????????????????
func isMatch(s string, p string) bool {
	// fmt.Printf("s: %s, p: %s\n", s, p)
	// sb, pb := []byte(s), []byte(p)
	p = removeDupAsteria(p)
	ls, lp := len(s), len(p)
	var curS, curP int // s,p??????????????????

	var lna int

	for i := 0; i < lp; i++ {
		if p[i] != asteria {
			lna++
		}
	}

	if lna > ls {
		// ??????????????????????????????????????????????????????????????????????????????
		return false
	}

	for {
		if curS >= ls {
			// s??????????????????
			if curP >= lp {
				// p???????????????
				return true
			}

			for i := curP; i < lp; i++ {
				if p[i] != asteria {
					return false // p??????????????????????????????????????????*???????????????
				}
			}
			return true
		}

		if curP >= lp {
			if curS >= ls {
				return true
			}
			return false
		}

		if p[curP] == asteria {
			// ????????????s???????????????p?????????????????????
			// ???????????????????????????
			// ????????????????????????????????????

			if curP == lp-1 {
				// '*'????????????????????????
				return true
			}

			var next []byte   // p??????????????????*?????????
			var nextIndex int // p??????????????????*????????????Index
			for i := curP; i < lp-1; i++ {
				b := p[i+1]
				if b != asteria && b != question {
					next = append(next, b)
					nextIndex = i + 1
				} else {
					if len(next) > 0 {
						// ??????????????????
						break
					}
				}
			}

			ln := len(next)
			if ln == 0 {
				// ??????????????????*???
				return true
			}

			var si []int // s??????next?????????
			for i := curS; i+ln <= ls; i++ {
				// if isMatch(s[i:i+ln], string(next)) {
				// 	si = append(si, i)
				// }

				if s[i:i+ln] == string(next) {
					si = append(si, i)
				}
			}

			lsi := len(si)
			if lsi == 0 {
				return false
			}

			// p1 := sub(p, 1)
			// p2 := sub(p, nextIndex+1)

			for i := lsi - 1; i >= 0; i-- {
				if ls-si[i] >= lna {
					res := isMatch(s[si[i]+ln:], p[1:])
					// res := isMatch(s[si[i]+ln:], p1)
					if res {
						return true
					}
					res = isMatch(s[si[i]+ln:], p[nextIndex+1:])
					// res = isMatch(s[si[i]+ln:], p2)
					if res {
						return true
					}

				}

			}
		}
		if p[curP] == question {
			curP++
			curS++ // match one
			continue
		}

		// ??????????????????
		if p[curP] != s[curS] {
			return false
		}
		curP++
		curS++
	}
}

func removeDupAsteria(p string) string {
	if len(p) == 0 {
		return p
	}
	// fmt.Println(p)
	result := []byte{p[0]}
	for i := 1; i < len(p); i++ {
		if p[i] != asteria || p[i-1] != asteria {
			result = append(result, p[i])
		}
	}
	return string(result)
}

func quickSort(src []int) {
	ls := len(src)
	if ls <= 1 {
		return
	}
	pivotIndex := 0
	pivot := src[pivotIndex]
	for i := 1; i < ls; i++ {
		if v := src[i]; v < pivot {
			copy(src[1:i+1], src[:i])
			src[0] = v
			pivotIndex++
		}
	}
	quickSort(src[:pivotIndex])
	quickSort(src[pivotIndex+1:])
}
