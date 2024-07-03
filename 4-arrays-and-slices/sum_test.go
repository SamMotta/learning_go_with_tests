package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run(
		"Collection of 5 numbers",
		func(t *testing.T) {
			// An Array!
			numbers := []int{1, 2, 3, 4, 5}

			got := Sum(numbers)

			expected := 15

			if got != expected {
				t.Errorf("Expected %d but got %d given, %v", expected, got, numbers)
			}
		},
	)

	t.Run(
		"Collection of any size",
		func(t *testing.T) {
			numbers := []int{1, 2, 3}

			got := Sum(numbers)

			expected := 6

			if got != expected {
				t.Errorf("Expected %d but got %d given, %v", expected, got, numbers)
			}
		},
	)
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v want %v", got, expected)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(got []int, expected []int, tb testing.TB) {
		tb.Helper()

		if !reflect.DeepEqual(got, expected) {
			tb.Errorf("got %v want %v", got, expected)
		}
	}

	t.Run(
		"Make the sum of some slices",
		func(t *testing.T) {
			got := SumAllTails([]int{1, 2}, []int{0, 9})
			expected := []int{2, 9}

			checkSums(got, expected, t)
		},
	)

	t.Run(
		"Safely sum empty slices",
		func(t *testing.T) {
			got := SumAllTails([]int{}, []int{0, 9})
			expected := []int{0, 9}

			checkSums(got, expected, t)
		},
	)
}
