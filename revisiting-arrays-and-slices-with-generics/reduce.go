package main

// Reduce function takes a slice of type A, an accumulator function, and an initial value of type A, and reduces the slice to a single value of type A.
func Reduce[A any](collection []A, accumulator func(A, A) A, initialValue A) A {
	var result = initialValue
	for _, x := range collection {
		result = accumulator(result, x)
	}
	return result
}

// Sum calculates the total from a slice of numbers using the Reduce function.
func Sum(numbers []int) int {
	add := func(acc, x int) int { return acc + x }
	return Reduce(numbers, add, 0)
}

// SumAllTails calculates the sums of all but the first number given a collection of slices using the Reduce function.
func SumAllTails(numbersToSum ...[]int) []int {
	sumTail := func(acc []int, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			tailSum := Sum(tail)
			return append(acc, tailSum)
		}
	}
	return Reduce(numbersToSum, sumTail, []int{})
}
