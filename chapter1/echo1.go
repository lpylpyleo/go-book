package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	for index, arg := range os.Args[1:] {
		fmt.Println(index, arg)
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
}
