package main

func Sum(numbers []int) (sum int) {
	reduce := func(a, b int) int {
		return a + b
	}

	return Reduce(numbers, reduce, 0)
}

func SumAllTails(numbers ...[]int) []int {
	reduce := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			return append(acc, Sum(x[1:]))
		}
	}
	return Reduce(numbers, reduce, []int{})
}

func Reduce[A any](collection []A, f func(A, A) A, initialValue A) A {
	result := initialValue
	for _, item := range collection {
		result = f(result, item)
	}

	return result
}
