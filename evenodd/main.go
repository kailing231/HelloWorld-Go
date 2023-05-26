package main

import "fmt"

func main() {
	numbers := []int{11, 22, 36, 40, 55, 69, 77, 88, 99, 103}

	for _, num := range numbers {
		if num%2 == 0 {
			fmt.Println(num, "even")
		} else {
			fmt.Println(num, "odd")
		}
	}
}
