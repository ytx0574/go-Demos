package utils

func FBN(n int) []int {
	if n <= 0 {
		return []int{}
	}

	var values []int = make([]int, n)
	if n >= 1 {
		values[0] = 1
	}
	if n >= 2 {
		values[1] = 1
	}
	if n >= 3 {
		for i := 2; i < n; i++ {
			values[i] = values[i - 1] + values[i - 2]
		}
	}
	return values
}

func AddUpper(n int) int {
	return n + 1
}