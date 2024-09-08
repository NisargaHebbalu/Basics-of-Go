package arrayandslices

import "fmt"

func main() {

	//arrays
	//var ages [3]int = [3]int{20, 25, 30}
	var ages = [3]int{20, 25, 30}

	names := [4]string{"John", "Jacob", "Melani", "Kyra"}
	names[1] = "Peter" //cannot append array

	//len(array_name) id used to find the length of the array
	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	//slices (use array under the hood)
	var scores = []int{100, 50, 60}
	scores[2] = 70
	scores = append(scores, 85) //can append in case of slice

	fmt.Println(scores, len(scores))

	//slice ranges - outer range is not included
	range1 := names[1:3]
	range2 := names[2:]
	range3 := names[:3]

	fmt.Println(range1, range2, range3)
	//fmt.Println(names[1:3])

	range1 = append(range1, "Kooper")

	fmt.Println(range1)

}
