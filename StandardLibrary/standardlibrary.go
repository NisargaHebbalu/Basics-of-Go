package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	greetings := "Hello, everyone!"

	//None of these change the actual string stored in the variable
	fmt.Println(strings.Contains(greetings, "Hello")) //It is case sensitive
	fmt.Println(strings.ReplaceAll(greetings, "Hello", "Hi"))
	fmt.Println(strings.ToUpper(greetings))
	fmt.Println(strings.Index(greetings, "ry"))
	fmt.Println(strings.Split(greetings, " "))

	fmt.Println("The original string:", greetings)

	//It changes original variable ages and keeps it sorted
	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30) //Returns the index number of that variable, if 90 is declared then the value returned will be one greater than the index number present in the ages
	fmt.Println(index)

	names := []string{"John", "Abele", "Kate", "Christian", "Mello"}
	sort.Strings(names)
	fmt.Println(names)

	stringindex := sort.SearchStrings(names, "Mello")
	fmt.Println(stringindex)
}

//Link for stardard library - https://www.youtube.com/redirect?event=video_description&redir_token=QUFFLUhqbXRENmxJdUxuSlNFZDQweWlscHdUbjRMM1NWd3xBQ3Jtc0trUXpmdU1pcF9MbDBVLTFVYWs5cGk1T0VCSzhZcEw0YVBya1A5MC1HVVpzS29BV0xweWZPQnlnQWQ2dWhzT1JzcWNSc1JDV0hTeXFQRlFodWxQOHY5enJRSjZ0QlMxLTFySTlLUXFwbkduV2EzWk5Jdw&q=https%3A%2F%2Fgolang.org%2Fpkg%2F&v=BoooHY57RXw
