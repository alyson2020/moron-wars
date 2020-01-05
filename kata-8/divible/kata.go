package kata

// IsDivisible checks n is divisible by x and y
func IsDivisible(n, x, y int) bool {
	divisibleByX := n%x == 0
	divisibleByY := n%y == 0
	// fmt.Printf("divisibleByX %v\n", divisibleByX)
	// fmt.Printf("divisibleByY %v\n", divisibleByY)
	return divisibleByX && divisibleByY
}
