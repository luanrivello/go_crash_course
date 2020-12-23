package main

import "fmt"

func main() {
	x := 9
	y := 10
	

	for i := 0; i < 3; i++ {
		if x < y {
			fmt.Printf("%d is less than %d\n", x, y)
		}else if x > y {
			fmt.Printf("%d is greater than %d\n", x, y)
		}else{
			fmt.Printf("%d is equal to %d\n", x, y)
		}

		x = x + 1

	}
	color := "red"

	switch color {
		case "red":
			fmt.Println("color is red")

		case "green":
			fmt.Println("color is green")

		default:
			fmt.Println("color is not green or red")

	}

}