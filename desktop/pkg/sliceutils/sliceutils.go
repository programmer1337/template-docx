package sliceutils

// Remove item from list by providing pos
func Remove(list []*int, i int) []*int {
	copy(list[i:], list[i+1:])
	list[len(list)-1] = nil
	list = list[:len(list)-1]

	return list
}

// Remove item from list by providing value
func RemoveByValue[T comparable](list []T, value T) []T {
	var defaultValue T

	for i := 0; i < len(list); i++ {
		if list[i] == value {
			copy(list[i:], list[i+1:])
			list[len(list)-1] = defaultValue
			list = list[:len(list)-1]
			i--
		}
	}

	return list
}

// Remove item from list by providing value
func RemoveByPValue[T comparable](list []*T, value *T) []*T {
	for i, v := range list {
		if v == value {
			copy(list[i:], list[i+1:])
			list[len(list)-1] = nil
			list = list[:len(list)-1]
		}
	}

	return list
}
