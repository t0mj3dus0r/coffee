package coffee

type List[E comparable] interface {
	Collection[E]

	// Returns the index of the first occurrence of the specified element in this list,
	// or -1 if this list does not contain the element.
	IndexOf(E) int
	// Removes the element at the specified position in this list.
	RemoveByIndex(index int) E
	// Replaces the element at the specified position in this list with the specified element.
	Set(index int, element E) E
	// Returns the element at the specified position in this list.
	Get(index int) E
}
