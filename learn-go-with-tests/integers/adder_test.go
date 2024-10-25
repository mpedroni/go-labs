package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

// it will be showed at `Add` function example section at godoc, out of the box.
// When running `go test`, it will be executed using the output value as expected value and fail if not match. That way
// we can trust our example code are always up to date and is compilable. For more information on testable examples, [see](https://go.dev/blog/examples)

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
