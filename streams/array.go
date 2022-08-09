package streams

import "github.com/t0mj3dus0r/coffee"

func FromArray[T any](array []T) coffee.Stream[T] {
	return arrayStream[T]{
		data: array,
	}
}

func Empty[T any]() coffee.Stream[T] {
	return arrayStream[T]{
		data: make([]T, 0),
	}
}

type arrayStream[T any] struct {
	data []T
}

// Returns whether all elements of this stream match the provided predicate.
func (s arrayStream[T]) AllMatch(predicate coffee.Predicate[T]) bool {

	for _, elem := range s.data {
		if !predicate.Test(elem) {
			return false
		}
	}

	return true
}

// Returns whether any elements of this stream match the provided predicate.
func (s arrayStream[T]) AnyMatch(predicate coffee.Predicate[T]) bool {

	for _, elem := range s.data {
		if predicate.Test(elem) {
			return true
		}
	}

	return false
}

// Creates a lazily concatenated stream whose elements are all the elements of the first stream followed by all the elements of the second stream.
func (s arrayStream[T]) Concat(stream coffee.Stream[T]) coffee.Stream[T] {

	var result []T

	result = append(result, s.data...)

	stream.ForEach(func(t T) {
		result = append(result, t)
	})

	return FromArray(result)
}

// Returns the count of elements in this stream.
func (s arrayStream[T]) Count() int {
	return len(s.data)
}

// Returns, if this stream is ordered, a stream consisting of the remaining elements of this stream after dropping the longest prefix of elements that match the given predicate.
func (s arrayStream[T]) DropWhile(predicate coffee.Predicate[T]) coffee.Stream[T] {
	panic("not implemented") // TODO: Implement
}

// Returns a stream consisting of the elements of this stream that match the given predicate.
func (s arrayStream[T]) Filter(predicate coffee.Predicate[T]) coffee.Stream[T] {

	var filtered []T

	for _, elem := range s.data {
		if predicate.Test(elem) {
			filtered = append(filtered, elem)
		}
	}

	return FromArray(filtered)
}

// Performs an action for each element of this stream.
func (s arrayStream[T]) ForEach(action func(T)) {
	for _, elem := range s.data {
		action(elem)
	}
}
