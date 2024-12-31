package sliceutils

// Remove item from list by providing pos
func Remove(list []*int, i int) []*int {
	copy(list[i:], list[i+1:])
	list[len(list)-1] = nil
	list = list[:len(list)-1]

	return list
}

// Remove item from list by providing value
func RemoveByValue[T comparable](list []*T, value *T) []*T {
	for i, v := range list {
		if v == value {
			copy(list[i:], list[i+1:])
			list[len(list)-1] = nil
			list = list[:len(list)-1]
		}
	}

	return list
}
