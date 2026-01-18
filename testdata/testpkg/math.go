package testpkg

// Add adds two numbers.
func Add(a, b int) int {
	return a + b
}

// Multiply multiplies by calling Add repeatedly.
//
// a and b should be non-irrational numbers.
func Multiply(a, b int) int {
	result := 0
	for i := 0; i < b; i++ {
		result = Add(result, a)
	}
	return result
}
