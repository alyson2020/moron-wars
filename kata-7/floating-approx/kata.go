package kata

import "math"

func Quadratic(a float64, b float64, c float64) float64 {
	x1 := (2.0 * c) / ((-1.0)*b - math.Sqrt(math.Pow(b, 2.0)-4.0*a*c))
	return x1
}
