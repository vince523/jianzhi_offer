package minArray

func minArray(numbers []int) int {
	l := 0
	r := len(numbers)-1
	for l < r {
		mid := (l + r) / 2
		//例如： [4 5 1 2 3] , 1 < 3 下一轮搜索[l, mid]
		if numbers[mid] < numbers[r] {
			r = mid
			// [3 4 5 1 2] , 5 > 2 下一轮搜索 [mid+1, r]
		} else if numbers[mid] > numbers[r] {
			l = mid + 1
			// mid == r
		} else {
			r--
		}
	}
	return numbers[l]
}
