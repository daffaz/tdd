package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Wawan")
	want := "Hello, Wawan"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
