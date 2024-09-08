package fmtmethods

import "fmt"

func main() {

	//fmt methods
	//Print
	fmt.Print("Hello, ")
	fmt.Print("world! \n")
	fmt.Print("Welcome to golang\n")

	//Println
	fmt.Println("Hello, world!")
	fmt.Println("Welcome back")

	name := "John"
	age := 26

	fmt.Println("My name is", name, "and my age is", age)

	//Printf(formatted strings)
	fmt.Printf("My name is %v and my age is %v\n", name, age) //%v is the default format specifier
	fmt.Printf("My name is %q and my age is %q\n", name, age) //%q for quotes but only for string
	fmt.Printf("Age is of type %T\n", age)                    //%T is used to specify the type of variable
	fmt.Printf("My score is %f\n", 24.55)                     //%f is used for floating point numbers
	fmt.Printf("My score is %0.2f\n", 24.55)                  //%0.2f - it rounds it upto 2 digit numbers after decimal, %0.1f - one digit after decimal

	//Sprintf(save formatted strings)
	var str = fmt.Sprintf("My name is %v and my age is %v\n", name, age)
	fmt.Println("Saved string is:", str)

}
