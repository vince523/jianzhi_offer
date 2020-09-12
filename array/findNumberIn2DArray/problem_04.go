package findNumberIn2DArray

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) < 1 {
		return false
	}
	n := len(matrix)-1
	m := 0
	for n >= 0 && m <= len(matrix[0])-1 {
		if matrix[n][m] > target {
			n--
		} else if matrix[n][m] < target{
			m++
		} else {
			return true
		}
	}
	return false
}
