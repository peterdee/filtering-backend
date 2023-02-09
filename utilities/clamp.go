package utilities

func Clamp(value, max, min int) int {
	if value > max {
		return max
	}
	if value < min {
		return min
	}
	return value
}
