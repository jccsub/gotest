package main

import (
	"fmt"
)

func main() {
	myName := "Chad Chandler"
	age := 48

	fmt.Println("\nTesting this for", myName, ". Is this working ", age)

	changeName(&myName)

	fmt.Println("Teting changeName method: ", myName)
}

func changeName(age *string) {
	*age = "new text"
}
