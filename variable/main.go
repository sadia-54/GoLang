package main

import 
(
	"fmt"
	"strings"
)

func main() {
	// var name string = "sadia"
	// var name2 = "jessia"
	// name3 := "farzana"
	// fmt.Println(name, name2, name3)
 
	// fmt.Println(num)

	// formatted string
	// fmt.Printf("hello, %v \n", name)

	// fmt.Printf("Score is %0.2f points", 200.222)

	var ages[3]int = [3]int{20, 30, 40}
	fmt.Println(ages)

	greeting := "hello there!"
	fmt.Println(strings.Contains(greeting, "there"))

		menu := map[string]float64{
		"burger": 5.99,
		"fries":  2.99,
		"soda":   1.49,
		"salad":  4.49,
	}

	fmt.Println(menu)
}