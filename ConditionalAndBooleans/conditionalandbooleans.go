package main

import (
	"fmt"
)

func main() {

	//True or False value
	age := 25

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("age is less than 30")
	} else if age < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is not less than 45")
	}

	names := []string{"John", "Abele", "Kate", "Christian", "Mello"}

	for indexnum, value := range names {
		if indexnum == 1 {
			fmt.Println("continuing at pos", indexnum)
			continue
		}

		if indexnum > 2 {
			fmt.Println("breaking at pos", indexnum)
			break
		}

		fmt.Printf("the value at pos %v is %v\n", indexnum, value)
	}
}
