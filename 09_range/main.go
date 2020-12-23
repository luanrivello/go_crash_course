package main

import "fmt"

func main() {
	ids := []int{55, 45, 35, 2}
	
	for i, id := range ids{
		fmt.Printf("%d-[%d]\n", i, id)
	}
	
	for _, id := range ids{
		fmt.Printf("%d\n", id)
	}
	
	sum := 0
	for _, id := range ids{
		sum = sum + id
	}

	fmt.Println("Sum =", sum)

	emails := map[string]string{"Bob": "bob@example.com", 
								"Jane": "jane@example.com", 
								"Tom": "tom@example.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}
	
	for k := range emails {
		fmt.Printf("Name: %s\n", k)
	}
}