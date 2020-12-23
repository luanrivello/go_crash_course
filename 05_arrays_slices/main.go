package main

import "fmt"

func main() {
	//Array
	//var fruitArr [2]string

	//Assgn value
	//fruitArr[0] = "apple"
	//fruitArr[1] = "orange"

	//Declare and assign
	fruitArr := [2]string{"apple", "orange"}

	fmt.Println(fruitArr)

	fruitSlice := []string{"apple", "orange", "banana", "bread"}
	fmt.Println(fruitSlice, len(fruitArr))
	fmt.Println(fruitSlice[1:3])

}