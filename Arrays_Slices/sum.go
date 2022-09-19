package arraysslices

func Sum(nums []int) (total int) {
	for _, num := range nums {
		total += num
	}
	return
}
