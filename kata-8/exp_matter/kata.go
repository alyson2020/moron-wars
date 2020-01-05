package kata

// ExpressionMatter returns int
func ExpressionMatter(a int, b int, c int) int {

	var result int

	for _, r := range calcExpressions(a, b, c) {
		// guarantees natural numbers
		if r > result {
			result = r
		}
	}
	return result
}

func calcExpressions(a, b, c int) []int {
	r1 := a + b + c
	r2 := a*b + c
	r3 := a + b*c
	r4 := a * b * c
	r5 := (a + b) * c
	r6 := a * (b + c)
	return []int{r1, r2, r3, r4, r5, r6}
}
