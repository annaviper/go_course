package main

import "fmt"

func main() {
	numbers := make([]int, 11)

	for i := range numbers {
		numbers[i] = i

		if i%2 == 0 {
			fmt.Printf("%v: even\n", i)
		} else {
			fmt.Printf("%v: odd\n", i)
		}

	}

	// fmt.Println(numbers)
}
