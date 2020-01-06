package kata

func century(year int) int {
	c := ((year - 1) / 100) + 1
	return c
}
