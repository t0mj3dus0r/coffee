package coffee

type Collection[E comparable] interface {
	Iterable[E]

	// Ensures that this collection contains the specified element.
	//
	// Returns true if this collection changed as a result of the call
	Add(e E) bool
	// Adds all of the elements in the specified collection to this collection.
	//
	// Returns true if this collection changed as a result of the call
	AddAll(Collection[E]) bool
	// Removes all of the elements from this collection.
	Clear()
	// Returns true if this collection contains the specified element.
	Contains(E) bool
	// Returns true if this collection contains all of the elements in the specified collection.
	ContainsAll(Collection[E]) bool
	// Returns true if this collection contains no elements.
	IsEmpty() bool
	// Returns an iterator over the elements in this collection.
	Iterator() Iterator[E]
	// Removes the specified element from this collection if it is present.
	//
	// Returns true if this collection contained the specified element
	Remove(E) bool
	// Removes all of the elements of this collection that satisfy the given predicate.
	//
	// Returns true if any elements were removed
	RemoveIf(filter Predicate[E]) bool
	// Removes from this collection all of its elements that are contained in the specified collection.
	//
	// Returns true if this collection changed as a result of the call
	RemoveAll(Collection[E]) bool
	// Returns the number of elements in this collection.
	Size() int
	// Returns a sequential Stream with this collection as its source.
	Stream() Stream[E]
	// Returns an array containing all of the elements in this collection.
	ToArray() []E
}
