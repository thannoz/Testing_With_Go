package arraysslices

import "testing"

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("want %d but got %d", want, got)
	}

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{12, 400, 3456, 9087654}

		got := Sum(numbers)
		want := 9091522

		if got != want {
			t.Errorf("want %d but got %d", want, got)
		}
	})
}

func BenchmarkSum(b *testing.B) {
	nums := []int{5, 4, 3, 2, 1}
	for i := 0; i < b.N; i++ {
		Sum(nums)
	}
}
