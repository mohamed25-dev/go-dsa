package array

func getConcatenation(nums []int) []int {
	for i := range len(nums) {
		nums = append(nums, nums[i])
	}

	return nums
}
