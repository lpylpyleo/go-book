package htmlout

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}

var res [][]int

func permute(nums []int) [][]int {
	root := &Node{999, []Node{}}
	add(root, nums)
	dump(nil, root)
	fmt.Println(res)

	return res
}

type Node struct {
	value    int
	children []Node
}

func add(n *Node, nums []int) []Node {
	var children []Node
	for i, num := range nums {
		cp := make([]int, len(nums))
		copy(cp, nums)
		cp[i] = cp[len(cp)-1]
		var c Node
		c = Node{num, add(&c, cp[:len(cp)-1])}
		children = append(children, c)
		n.children = append(n.children, c)
	}
	return children
}

func dump(seen []int, n *Node) {
	seen = append(seen, n.value)
	if len(n.children) == 0 {
		res = append(res, seen[1:])
	}

	for _, c := range n.children {
		seenCpy := make([]int, len(seen))
		copy(seenCpy, seen)
		dump(seenCpy, &c)
	}
}
