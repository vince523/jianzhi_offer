package findRepeatNumber

import "sort"

func findRepeatNumber(nums []int) int {
	// 原地置换
	for i, v := range nums {
		if nums[i] == i {
			// 位置放的值正确，跳过
			continue
		}
		if nums[i] == nums[nums[i]] {
			// 说明重复了
			return v
		}

		nums[i], nums[v] = nums[v], nums[i]
	}
	return 0
}

// 解法2 map
func findRepeatNumber2(nums []int) int {
	tmp := make(map[int]int, 0)
	for i, v := range nums{
		if _, ok := tmp[v]; ok {
			return v
		}
		tmp[v] = i
	}
	return 0
}

// 排序
func findRepeatNumber3(nums []int) int {
	sort.Ints(nums)
	for i, v := range nums {
		if v == nums[i+1] {
			return v
		}
	}
	return 0
}