package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	// a := [3]int{1, 2, 3}
	// fmt.Println(a[0])
	// fmt.Println(a[len(a)-1])

	// b := [...]rune{99: 'a'}
	// fmt.Println(b[len(b)-1])

	// c1 := sha256.Sum256([]byte{'x', 'y'})
	// c2 := sha256.Sum256([]byte{'X', 'y'})
	// fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	// fmt.Println(diffBitCounts(c1, c2))

	writeHash()
}

func diffBitCounts(a, b [32]byte) int {
	var count int
	for i := 0; i < 32; i++ {
		xor := a[i] ^ b[i]
		count += popcount(int(xor))
	}
	return count
}

func popcount(x int) int {
	var count int

	for x > 0 {
		x = x & (x - 1)
		count++
	}

	return count
}

var algorithm = flag.String("a", "sha256", "The algorithm used, sha256, sha384 or sha512 available.")

func writeHash() {
	input := bufio.NewScanner(os.Stdin)
	flag.Parse()
	a := *algorithm
	input.Scan()
	if a == "sha256" {
		result := sha256.Sum256(input.Bytes())
		fmt.Printf("sha256: %x\n", result)
	} else if a == "sha384" {
		result := sha512.Sum384(input.Bytes())
		fmt.Printf("sha384: %x\n", result)
	} else if a == "sha512" {
		result := sha512.Sum512(input.Bytes())
		fmt.Printf("sha512: %x\n", result)
	} else {
		os.Exit(1)
	}
}
