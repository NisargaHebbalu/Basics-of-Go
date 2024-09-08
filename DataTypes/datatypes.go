package datatypes

import "fmt"

func main() {

	fmt.Println("Hello, ninjas!")

	//strings
	//interpreted string literals
	var name1 string = "Nisarga"
	var name2 = "H N"
	var name3 string

	fmt.Println(name1, name2, name3)

	name1 = "Ni"
	name3 = "Nisi"

	fmt.Println(name1, name2, name3)

	name4 := "Nisu" //we can't use this outside the function

	fmt.Println(name4)

	name4 = "Chittu"

	fmt.Println(name4)

	//raw string literals
	var sentence string = `Mary had a little lamb,
little lamb, little lamb,
Mary had a little lamb,
My fair lady`

	fmt.Println(sentence)

	//ints
	var num1 int = 25
	var num2 = 26
	var num3 int
	num3 = 27
	num4 := 28

	fmt.Println(num1, num2, num3, num4)

	//int8(-128 to 127), int16, innt32, int64 differce in range
	var number1 int8 = 25
	var number2 uint = 78 //unsigned number which means we cannot have negative number
	//uint can also have the bit size - uint8(but here we can go beyond 127 since we cannot have negative numbers, range =(0 to 255)), uint16...

	fmt.Println(number1, number2)

	//floats - unlike integers we will have to assign the bit size as either 32 or 64
	var score1 float32 = -64.66
	var score2 float64 = 54.8974562977
	score3 := 1.5 //by default it is float64

	fmt.Println(score1, score2, score3)

}
