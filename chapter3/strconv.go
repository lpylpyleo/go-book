package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 123
	x := fmt.Sprintf("%8b\n", a)
	fmt.Println(x, strconv.FormatInt(int64(a), 2))
}
