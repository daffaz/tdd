package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Wawan", "")
		want := "Hello, Wawan"
		assertCorrectMessage(t, got, want)
	})
	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Wawan", "es")
		want := "Hola, Wawan"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in france", func(t *testing.T) {
		got := Hello("Wawan", "fr")
		want := "Bonjour, Wawan"
		assertCorrectMessage(t, got, want)
	})
}
