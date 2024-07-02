package main

import "testing"

func TestHello(t *testing.T) {

	t.Run(
		"Saying hello to people",
		func(t *testing.T) {
			got := Hello("Sam", "")
			want := "Hello, Sam"

			assertCorrectMessage(t, got, want)
		},
	)

	t.Run(
		"Empty string defaults to 'World'",
		func(t *testing.T) {
			got := Hello("", "")
			want := "Hello, World"

			assertCorrectMessage(t, got, want)
		},
	)

	t.Run(
		"In spanish",
		func(t *testing.T) {
			got := Hello("Elodie", "Spanish")
			want := "Hola, Elodie"

			assertCorrectMessage(t, got, want)
		},
	)

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("Got %q, but want %q", got, want)
	}
}
