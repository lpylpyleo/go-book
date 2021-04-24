package main

import (
	"fmt"
	"time"
)

func main() {
	// var dilbert employee
	// fmt.Printf("%T %[1]v\n", dilbert)
	// s := []int{54, 32, 64, 657, 23, 47, 235}
	// Sort(s)
	// fmt.Println(s)

	wheel := Wheel{spokes: 10}
	wheel.x = 123
	fmt.Printf("%#v\n", wheel)
}

type employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

// ---------------------------------------
type Point struct {
	x, y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	spokes int
}
