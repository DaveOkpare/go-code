package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(1, 2)
	expected := 3

	if sum != expected {
		t.Errorf("expect '%d', got '%d'", expected, sum)
	}

}

func ExampleAdd() {
	sum := Add(1, 2)
	fmt.Println(sum)
	// Output: 3
}
