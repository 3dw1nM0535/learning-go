package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		// testing.TB satisfy test and benchmark interface so you
		// call helper functions from both
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("should say hello to people", func(t *testing.T) {
		got := Hello("Edwin")
		want := "Hello, Edwin"

		assertCorrectMessage(t, got, want)
	})

	t.Run("should say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

}
