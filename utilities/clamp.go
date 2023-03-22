package utilities

func Clamp[T float64 | int | uint](value, max, min T) T {
	if value > max {
		return max
	}
	if value < min {
		return min
	}
	return value
}
