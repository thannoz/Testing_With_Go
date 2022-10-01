package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Carlos")

	got := buffer.String()
	want := "Hello, Carlos"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
