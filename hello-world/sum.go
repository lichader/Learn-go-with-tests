package main

import "fmt"

func Sum(numbers []int) (sum int) {
	for index, number := range numbers {
		fmt.Printf("index: %d, number: %d\n", index, number)
		sum += number
	}
	return
}

func SumAllTails(numbers ...[]int) []int {
	var sums []int
	for _, number := range numbers {

		if len(number) == 0 {
			sums = append(sums, 0)
			continue
		}

		tail := number[1:]
		sums = append(sums, Sum(tail))
	}
	return sums
}
