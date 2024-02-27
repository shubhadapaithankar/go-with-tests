package iteration

import "testing"

// BenchmarkRepeat benchmarks the Repeat function.
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a") // Directly use Repeat from iteration.go
	}
}

// TestRepeat tests the Repeat function.
func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
