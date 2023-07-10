package run

func MapCopy(m map[string]float64) map[string]float64 {
	result := make(map[string]float64)
	for k, v := range m {
		result[k] = v
	}
	return result
}

func IntMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
