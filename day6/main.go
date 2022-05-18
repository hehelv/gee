package main

import "fmt"

func main() {
	test := []int{1}
	test = append(test, 2)

	fmt.Println(test[1])
}
