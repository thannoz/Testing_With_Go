package arraysslices

func Sum(nums []int) (total int) {
	for _, num := range nums {
		total += num
	}
	return
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNums := len(numbersToSum)
	sums := make([]int, lengthOfNums)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
