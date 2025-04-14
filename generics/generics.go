package generics


func Reverse[T any](input []T) []T {
	n := len(input)

	result := make([]T, n)
	for i,v := range input {
		result[n-1-i] = v
	}
	
	return result
}