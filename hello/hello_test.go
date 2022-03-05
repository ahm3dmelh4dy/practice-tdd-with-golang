package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Ahmed")
	want := "Hello, Ahmed"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
