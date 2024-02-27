// hello_test.go
package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to someone", func(t *testing.T) {
		got := Hello("David", "")
		want := "Hello, David!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello when no name is given", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("David", "Spanish")
		want := "Hola, David!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("David", "French")
		want := "Bonjour, David!"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
