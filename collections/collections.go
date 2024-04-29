package collections

func Contains[Item comparable](items []Item, searchValue Item) bool {
	for _, item := range items {
		if item == searchValue {
			return true
		}
	}

	return false
}

func MapSlice[Input any, Output any](input []Input, transformer func(Input) Output) []Output {
	var result []Output
	for _, item := range input {
		result = append(result, transformer(item))
	}

	return result
}
