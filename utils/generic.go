package Generic

func ConvertArrayToArrayOfPointers[T any](x []T) []*T {
	// Function body using type T
	var result []*T
	for _, item := range x {
		result = append(result, &item)
	}

	return result
}
