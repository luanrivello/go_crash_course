package main

import "fmt"

func main() {
	sum := adder(100)

	for i := 0; i < 10; i++ {
		fmt.Println(sum(1))
	}

}

func adder(a int) func(int) int {
	var sum = a

	return func(x int) int{
		sum = sum + x
		return sum
	}
}