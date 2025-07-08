package array

func removeElement(nums []int, val int) int {
	l := 0

	for _, num := range nums {
		if num == val {
			continue
		}

		nums[l] = num
		l++
	}

	return l
}
