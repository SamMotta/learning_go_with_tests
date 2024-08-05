package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)

		expected := 15

		if got != expected {
			t.Errorf("Expected %d but got %d given, %v", expected, got, numbers)
		}
	})

	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)

		expected := 6

		if got != expected {
			t.Errorf("Expected %d but got %d given, %v", expected, got, numbers)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("Make the sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}

		checkSums(got, expected, t)
	})

	t.Run("Safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		expected := []int{0, 9}

		checkSums(got, expected, t)
	})
}

func checkSums(got []int, expected []int, tb testing.TB) {
	tb.Helper()

	if !reflect.DeepEqual(got, expected) {
		tb.Errorf("got %v want %v", got, expected)
	}
}

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}

		AssertEqual(t, Reduce([]int{1, 2, 3}, multiply, 1), 6)
	})

	t.Run("concatenate strings", func(t *testing.T) {
		concatenate := func(x, y string) string {
			return x + y
		}

		AssertEqual(t, Reduce([]string{"a", "b", "c"}, concatenate, ""), "abc")
	})
}

func TestFind(t *testing.T) {
	t.Run("find first even number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		firstEvenNumber, found := Find(numbers, func(x int) bool { return x%2 == 0 })

		AssertTrue(t, found)
		AssertEqual(t, firstEvenNumber, 2)
	})
}
