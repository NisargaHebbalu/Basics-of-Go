package main

import (
	"fmt"
)

func main() {

	x := 0
	for x < 5 {
		fmt.Println("value of x is:", x)
		x++
	}

	for i := 0; i < 5; i++ {
		fmt.Println("value of i is:", i)
	}

	names := []string{"John", "Abele", "Kate", "Christian", "Mello"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, value := range names {
		fmt.Printf("the value at index %v is %v \n", index, value)
	}

	for _, value := range names { // If we don't want to use either index or the value of the slice or array just put an underscore
		fmt.Printf("the value is %v \n", value)
		value = "new string" // this doesn't update the value of the slice bcoz it acts like a local variable
	}

	fmt.Println(names)

}
