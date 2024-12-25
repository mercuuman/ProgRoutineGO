package mathutils

func Factorial(a int) int {
	if a == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= a; i++ {
		result *= i
	}
	return result
}
