package sliceutils

type Array[T comparable] []T

// Returns the number of elements in Array
func (arr Array[T]) Length() int {
	return len(arr)
}

// Push appends elements to the end of the array.
func (arr *Array[T]) Push(elements ...T) {
	*arr = append(*arr, elements...)
}

// Pop removes the last element from an array and returns that element.
// This method changes the length of the array.
func (arr *Array[T]) Pop() T {
	lastIdx := len(*arr) - 1
	lastEl := (*arr)[lastIdx]
	*arr = (*arr)[:lastIdx]

	return lastEl
}

// Executes a provided function once for each array element.
func (arr Array[T]) ForEach(iterator func(element T, index int, arr Array[T])) {
	for idx, el := range arr {
		iterator(el, idx, arr)
	}
}

// Creates a new array populated with the results of
// calling a provided function on every element in the calling array.
func (arr Array[T]) Map(iterator func(element T, index int, arr Array[T]) T) Array[T] {
	newArr := make([]T, arr.Length())

	for idx, el := range arr {
		newArr[idx] = iterator(el, idx, arr)
	}

	return newArr
}

// Filter creates a new array with all elements that pass the test implemented by the provided function.
func (arr Array[T]) Filter(iterator func(element T, index int, arr Array[T]) bool) Array[T] {
	var newArr []T

	for idx, el := range arr {
		if iterator(el, idx, arr) {
			newArr = append(newArr, el)
		}
	}

	return newArr
}

// IndexOf returns the first index at which a given element can be found in the array, or -1 if it is not present.
func (arr Array[T]) IndexOf(element T) int {
	for idx, el := range arr {
		if el == element {
			return idx
		}
	}

	return -1
}

// Includes determines whether an array includes a certain element, returning true or false as appropriate.
func (arr Array[T]) Includes(element T) bool {
	return arr.IndexOf(element) != -1
}
