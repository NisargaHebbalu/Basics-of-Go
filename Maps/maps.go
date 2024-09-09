package main

import (
	"fmt"
)

func main() {

	//unlike python a key can be of any type
	//but in a variable key must follow the same type and value must also be of the same type

	// all the keys must be string and values must be int
	menu := map[string]int{
		"Soup":  120,
		"Salad": 140,
		"Meal":  210,
		"Juice": 50,
	}

	fmt.Println(menu)
	fmt.Println(menu["Meal"])

	//looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	//keys as int
	phonebook := map[int]string{
		9010101010: "Priya",
		9020202020: "Ritu",
		9030303030: "Abhi",
		9040404040: "Arya",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[9010101010])

	for k, v := range phonebook {
		fmt.Println(k, "-", v)
	}

	//updating the map
	phonebook[9050505050] = "Avni"
	fmt.Println(phonebook)

	phonebook[9040404040] = "Arush"
	fmt.Println(phonebook)

}
