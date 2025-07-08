package array

func removeDuplicates(arr []int) int {
	n := len(arr)
	r := 0
	l := 0

	for r < n {
		arr[l] = arr[r]
		for r < n && arr[l] == arr[r] {
			r++
		}
		l++
	}

	return l
}
