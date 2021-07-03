package utils

// FindElement returns a index of a element (e)
func FindElement(elements []string, e string) int {
	for i := range elements {
		if elements[i] == e {
			return i
		}
	}
	return -1
}

// RemoveElement removes the element passed from the slice
func RemoveElement(elements []string, e string) []string {
	i := FindElement(elements, e)
	return append(elements[:i], elements[i+1:]...)
}
