package main

import "fmt"

func main() {
	const isCool = true
	//var name = "Luan"
	//var age = 246

	name := "Luan"
	age := 246

	if name == "Luan" {
		age++
	}

	fmt.Println("----------------------------------------------------------------")
	fmt.Println(name, age, isCool)
	fmt.Printf("%T %T %T\n", name, age, isCool)
	fmt.Println("----------------------------------------------------------------")

}
