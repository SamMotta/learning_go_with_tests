package main

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	repeated := Repeat("A", 3)

	fmt.Println(repeated) // Output: AAA
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("Expected '%q' but got '%q'", expected, repeated)
	}

}

func BenchmarkRepeat(b *testing.B) {
	for range b.N {
		Repeat("a", 5)
	}
}
