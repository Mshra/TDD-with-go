package iteration

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	repeat := Repeat("a", 8)
	fmt.Println(repeat)
	// output: aaaaaaaa
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 10)
	expected := "aaaaaaaaaa"

	if repeated != expected {
		t.Errorf("exptected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
