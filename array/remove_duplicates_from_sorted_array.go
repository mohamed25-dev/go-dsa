package array

func removeDuplicatesOption1(arr []int) int {
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

func removeDuplicatesOption2(arr []int) int {
	l := 1
	for r := 1; r < len(arr); r++ {
		if arr[r] != arr[r-1] {
			arr[l] = arr[r]
			l++
		}
	}

	return l
}
