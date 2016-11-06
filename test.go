package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	defer fmt.Println("End!")
	var i, k int
	i = 5
	k = 6
	x := add(i, k)
	fmt.Println(swap(x, k))
	var l, m = true, "Hi!"
	fmt.Println(l, m)
	fmt.Println(multi(4))
	const format = "%T(%v)\n"
	o := 5
	p := float64(o)
	fmt.Printf(format, o, o)
	fmt.Printf(format, p, p)
	q := 1
	for q < 1000 {
		q += q;
	}
	fmt.Println(q)
	fmt.Println(max(q, x))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	week()
	say()
}

func add(x, y int) int {
	return x + y
}

func swap(x, y int) (int, int) {
	return y, x
}

func multi(x int) (sum int) {
	sum = 1
	for i := 1; i <= x; i++ {
		sum *= i;
	}
	return
}

func max(x, y int) (sum int) {
	sum = x
	if x < y {
		sum = y
	}
	return
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func week() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func say() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}