package spiralOrder

func spiralOrder(matrix [][]int) []int {
	if len(matrix) < 1 {
		return []int{}
	}
	var (
		left = 0
		right = len(matrix[0])-1
		top = 0
		bottom = len(matrix)-1
		result = make([]int, 0)
	)

	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}

		for j := top+1; j <= bottom; j++ {
			result = append(result, matrix[j][right])
		}

		if left < right && top < bottom {
			for q := right-1; q > left; q-- {
				result = append(result, matrix[bottom][q])
			}
			for p := bottom; p > top; p-- {
				result = append(result, matrix[p][left])
			}
		}
		left++
		right--
		top++
		bottom--

	}

	return result
}
