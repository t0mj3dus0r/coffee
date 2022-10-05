package coffee

type Stream[T any] interface {
	// Returns whether all elements of this stream match the provided predicate.
	AllMatch(predicate Predicate[T]) bool
	// Returns whether any elements of this stream match the provided predicate.
	AnyMatch(predicate Predicate[T]) bool
	// Creates a lazily concatenated stream whose elements are all the elements of the first stream followed by all the elements of the second stream.
	Concat(stream Stream[T]) Stream[T]
	// Returns the count of elements in this stream.
	Count() int
	// Returns, if this stream is ordered, a stream consisting of the remaining elements of this stream after dropping the longest prefix of elements that match the given predicate.
	DropWhile(predicate Predicate[T]) Stream[T]
	// Returns a stream consisting of the elements of this stream that match the given predicate.
	Filter(predicate Predicate[T]) Stream[T]
	// Performs an action for each element of this stream.
	ForEach(action func(T))

	Map(mapFn func(T) T) Stream[T]

	ToList() []T
}
