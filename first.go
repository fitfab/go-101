package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func add(x, y int) int {
	return x + y
}

func swap(w, u string) (string, string) {
	return u, w
}

func printThis() {
	var name, phone, city = "Miguel", "917-009-3311", "New York"
	fmt.Printf("name: %v phone: %v city: %q\n", name, phone, city)
}

func loopThis() {
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func whileThis() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)
}

func ifThis(foo int64) string {
	if foo > 100 {
		return fmt.Sprint(foo, " is greater than 100")
	}
	return fmt.Sprint("this is foo: ", foo)
}

func switchThis() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon -- almost done.")
	default:
		fmt.Println("Good evening!")
	}
}

func deferThis() {
	defer fmt.Println("this is called last")

	fmt.Println("this is the call at end of function")
}

func main() {
	os := runtime.GOOS

	fmt.Printf("you are running %v os\n", os)
	k := 100
	var i, c int = 2, 3
	var python, java bool
	fmt.Println("1st Hello World!!!", math.Sqrt(7))

	fmt.Println(add(4, 8))

	a, b := swap("world!", "Hello")
	fmt.Println("swap this 'world! Hello' ", a, b)

	fmt.Println("The time is", time.Now())

	fmt.Println(i, c, python, java, k)

	printThis()

	loopThis()

	whileThis()

	fmt.Println(ifThis(120))

	switchThis()

	deferThis()
}
