package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Sam")

	got := buffer.String()
	expected := "Hello, Sam"

	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}
}
