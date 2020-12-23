package main

import "fmt"

func main() {
	//map
	//emails := make(map[string]string)

	//emails["Bob"] = "bob@example.com"
	//emails["Jane"] = "jane@example.com"
	//emails["Tom"] = "tom@example.com"

	emails := map[string]string{"Bob": "bob@example.com", 
								"Jane": "jane@example.com", 
								"Tom": "tom@example.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	delete(emails, "Bob")

	fmt.Println(emails)

}