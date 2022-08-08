package coffee

type Set[T comparable] interface {
	// Adds the specified element to this set if it is not already present.
	// Returns true if this set did not already contain the specified element
	Add(T) bool
	// Adds all of the elements in the specified collection to this set if they're not already present.
	// Returns true if this set changed as a result of the call
	AddAll([]T) bool
	// Removes all of the elements from this set.
	Clear()
	// Returns true if this set contains the specified element.
	Contains(T) bool
	// Returns true if this set contains all of the elements of the specified collection.
	ContainsAll([]T) bool
	// Returns true if this set contains no elements.
	IsEmpty() bool
	// Returns the number of elements in this set (its cardinality).
	Size() int
	// Returns an array containing all of the elements in this set.
	ToArray() []T
	// Removes the specified element from this set if it is present.
	// Returns true if this set contained the specified element
	Remove(T) bool
	// Removes from this set all of its elements that are contained in the specified collection.
	// Returns true if this set changed as a result of the call
	RemoveAll([]T) bool
	// Retains only the elements in this set that are contained in the specified collection.
	// Returns true if this set changed as a result of the call
	RetainAll([]T) bool
}
