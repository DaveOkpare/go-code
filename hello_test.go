// hello_test.go
package main // <-- Importing the right package here.

import "testing"

func TestHello(t *testing.T) {
	got := Hello("David")
	want := "Hello, David!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
