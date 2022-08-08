package set

func NewHashSet[T comparable]() HashSet[T] {
	return HashSet[T]{
		data: make(map[T]struct{}),
	}
}

type HashSet[T comparable] struct {
	data map[T]struct{}
}

func (h HashSet[T]) Add(t T) bool {
	if h.Contains(t) {
		return false
	}

	h.data[t] = struct{}{}

	return true
}

func (h HashSet[T]) AddAll(t []T) bool {

	var result bool

	for _, elem := range t {
		setModified := h.Add(elem)
		if setModified {
			result = true
		}
	}

	return result
}

func (h HashSet[T]) Clear() {
	// Optimized by Go compiler. This snippet avoids eventual memory leaks / garbage collecting
	for k := range h.data {
		delete(h.data, k)
	}
}

func (h HashSet[T]) Contains(t T) bool {
	_, exists := h.data[t]

	return exists
}

func (h HashSet[T]) ContainsAll(t []T) bool {

	for _, elem := range t {
		if !h.Contains(elem) {
			return false
		}
	}

	return true
}

func (h HashSet[T]) IsEmpty() bool {
	return len(h.data) == 0
}

func (h HashSet[T]) Size() int {
	return len(h.data)
}

func (h HashSet[T]) ToArray() []T {

	result := make([]T, len(h.data))

	for k := range h.data {
		result = append(result, k)
	}

	return result
}

func (h HashSet[T]) Remove(t T) bool {
	if !h.Contains(t) {
		return false
	}

	delete(h.data, t)

	return true
}

func (h HashSet[T]) RemoveAll(t []T) bool {

	var result bool

	for _, elem := range t {
		setModified := h.Remove(elem)
		if setModified {
			result = true
		}
	}

	return result
}

func (h HashSet[T]) RetainAll(t []T) bool {

	existingElements := make([]T, len(t))

	for _, elem := range t {
		if h.Contains(elem) {
			existingElements = append(existingElements, elem)
		}
	}

	result := len(existingElements) != len(h.data)

	h.Clear()

	h.AddAll(existingElements)

	return result
}
