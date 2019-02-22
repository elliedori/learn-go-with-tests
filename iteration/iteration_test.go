package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 12)
	expected := "aaaaaaaaaaaa"

	if repeated != expected {
		t.Errorf("expected '%s' but got '%s'", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 12)
	}
}

func ExampleRepeat() {
	repeated := Repeat("j", 7)
	fmt.Println(repeated)
	// Output: jjjjjjj
}
