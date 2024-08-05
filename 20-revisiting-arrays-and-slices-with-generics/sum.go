package main

func Sum(numbers []int) int {
	add := func(acc, v int) int { return acc + v }
	return Reduce(numbers, add, 0)
}

func SumAllTails(numbersToSum ...[]int) []int {
	var zero []int

	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, Sum(tail))
		}
	}

	return Reduce(numbersToSum, sumTail, zero)
}
