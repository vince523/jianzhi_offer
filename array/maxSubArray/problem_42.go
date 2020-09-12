package maxSubArray

func maxSubArray(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	sum := nums[0]


	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		if dp[i] > sum {
			sum = dp[i]
		}
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
