package coffee

type Iterator[E any] interface {
	// Performs the given action for each remaining element until
	// all elements have been processed or the action throws an exception.
	ForEachRemaining(action Consumer[E])
	// Returns true if the iteration has more elements.
	HasNext() bool
	// Returns the next element in the iteration.
	Next() E
}
