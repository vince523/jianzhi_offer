package findContinuousSequence

func findContinuousSequence(target int) [][]int {
	// 双指针
	l, r := 1, 2
	ret := make([][]int, 0)
	for l < r {
		// 首项+末项 * 项数 / 2
		sum := (l + r) * (r-l+1) / 2
		if sum == target {
			tmp := make([]int, 0)
			for i := l; i <= r; i++ {
				tmp = append(tmp, i)
			}
			ret = append(ret, tmp)
			// 右移 左指针
			l++
		} else if sum < target {
			r++
		} else {
			l++
		}
	}
	return ret
}
