package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 10)
	expected := "aaaaaaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestStringsRepeat(t *testing.T) {
	repeated := strings.Repeat("x", 7)
	expected := "xxxxxxx"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

// run the func N times and print results. Run with `go test -bench=.`, where .
// N is defined by Go internally
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("x", 6)
	fmt.Println(repeated)

	// Output: xxxxxx
}
