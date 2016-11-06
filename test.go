package main

import (
	"fmt"
	"math"
	"time"
)

type Com struct {
	a int
	i int
}

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
	pointer()
	com := Com{1, 2}
	fmt.Println(com)
	com1 := &com
	fmt.Println(com1.a)
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(s)
	s = s[1:4]
	fmt.Println(s)
	s = s[:2]
	fmt.Println(s)
	s = s[1:]
	fmt.Println(s)
	var s1 []int
	fmt.Println(s1, len(s1), cap(s1))
	if s1 == nil {
		fmt.Println("nil!")
	}
	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)

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

func pointer() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}