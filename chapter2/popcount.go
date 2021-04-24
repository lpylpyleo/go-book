package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	// for _, v := range pc {
	// 	fmt.Printf("%b\n", v)
	// }

	start := time.Now()
	for _, arg := range os.Args[1:] {
		num, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("%s is not a valid integer.\n", arg)
			continue
		}
		fmt.Printf("Pop count of %d is %d.\n", num, Popcount(uint64(num)))
	}
	fmt.Printf("Pop1: %s elapsed.\n", time.Since(start))

	start = time.Now()
	for _, arg := range os.Args[1:] {
		num, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("%s is not a valid integer.\n", arg)
			continue
		}
		fmt.Printf("Pop count of %d is %d.\n", num, Popcount2(uint64(num)))
	}
	fmt.Printf("Pop2: %s elapsed.\n", time.Since(start))

	start = time.Now()
	for _, arg := range os.Args[1:] {
		num, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("%s is not a valid integer.\n", arg)
			continue
		}
		fmt.Printf("Pop count of %d is %d.\n", num, Popcount3(uint64(num)))
	}
	fmt.Printf("Pop3: %s elapsed.\n", time.Since(start))
}

func Popcount(x uint64) int {
	return int(
		pc[byte(x>>(0*8))] +
			pc[byte(x>>(1*8))] +
			pc[byte(x>>(2*8))] +
			pc[byte(x>>(3*8))] +
			pc[byte(x>>(4*8))] +
			pc[byte(x>>(5*8))] +
			pc[byte(x>>(6*8))] +
			pc[byte(x>>(7*8))],
	)
}

func Popcount2(x uint64) int {
	var count int

	for i := 0; i < 64; i++ {
		d := (x >> i) & 1
		if d > 0 {
			count++
		}
	}

	return count
}

func Popcount3(x uint64) int {
	var count int

	for ; x > 0; count++ {
		x = x & (x - 1)
	}

	return count
}
