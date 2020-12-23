package main

import (
	"fmt"
)

func main() {
	var result [2]int
	result[0], result[1]= getSum(1,2)
	fmt.Println(greetings("traveller"), result[0], result[1])

}

func greetings(name string) string {
	return "Greetings " + name
}

func getSum(num1 int, num2 int) (int,int) {
	return num1, num2
}
