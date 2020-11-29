package main

import "fmt"

func assignment() {
	N := 11
	numbers := make([]int, N)
	for i := 0; i < N; i++ {
		numbers[i] = i
	}

	for i := range numbers {
		evenOrOdd := ""
		if i%2 == 0 {
			evenOrOdd = "even"
		} else {
			evenOrOdd = "odd"
		}
		fmt.Println(i, " is ", evenOrOdd)
	}
}
