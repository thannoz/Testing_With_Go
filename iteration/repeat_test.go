package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 4)
	exptected := "aaaa"

	if repeated != exptected {
		t.Errorf("expected %q but got %q", exptected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 1)
	}
}
