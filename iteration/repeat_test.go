package iteration

import (
	"fmt"
	"testing"
)

func TestIterate(t *testing.T) {
	got := Iterate("a", 5)
	expected := "aaaaa"

	if got != expected {
		t.Errorf("got %q, expected %q", got, expected)
	}
}

func BenchmarkIterate(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Iterate("a", 5)
	}
}

func ExampleIterate() {
	sequence := Iterate("a", 9)
	fmt.Println(sequence)
	// Output: aaaaaaaaa
}
