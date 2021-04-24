package main

import (
	"fmt"
	"time"
)

const (
	// sl = [1]int{}
	length = len("123")
	// capacity = cap( )
	r = real(2 + 3i)
	i = imag(2 + 3i)
	c = complex(2, 3)
)

func main() {
	// showType()
	// iotaExample()
	// netFlag()
	units()
}

//------------------------------------------

const (
	KB = 1000
	MB = 1000 * KB
	GB = 1000 * MB
	TB = 1000 * GB
	PB = 1000 * TB
	EB = 1000 * PB
	ZB = 1000 * EB
	YB = 1000 * ZB
)

func units() {
	fmt.Printf("%e\n", KB/1.0)
	fmt.Printf("%e\n", MB/1.0)
	fmt.Printf("%d\n", YB/ZB)
}

//------------------------------------------

type Flags uint

const (
	FlagUp Flags = 1 << iota
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

func netFlag() {
	v := FlagUp | FlagMulticast
	fmt.Printf("Is up? %v\n", v&FlagUp == FlagUp)
	fmt.Printf("Is loopback? %v\n", v&FlagLoopback == FlagLoopback)
	fmt.Printf("Is multicast? %v\n", v&FlagMulticast == FlagMulticast)
}

//------------------------------------------

func showType() {
	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T\t%[1]v\n", noDelay)
	fmt.Printf("%T\t%[1]v\n", timeout)
	fmt.Printf("%T\t%[1]v\n", time.Minute)
}

// --------------------------------------

type Weekday int

func (d Weekday) String() string {
	return fmt.Sprintf("星期%d\n", d)
}

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func iotaExample() {
	fmt.Println(Monday)
}
