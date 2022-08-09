package streams

import "github.com/t0mj3dus0r/coffee"

// Returns a stream consisting of the results of applying the given function to the elements of the stream.
func Map[T, V any](stream coffee.Stream[T], mapper func(T) V) coffee.Stream[V] {

	var result []V

	stream.ForEach(func(t T) {
		result = append(result, mapper(t))
	})

	return FromArray(result)
}

// Performs a mutable reduction operation on the elements of the stream using a Collector.
func Collect[T, A, R any](stream coffee.Stream[T], collector coffee.Collector[T, A, R]) R {

	container := collector.Supplier()()
	accumulator := collector.Accumulator()

	stream.ForEach(func(t T) {
		accumulator(container, t)
	})

	return collector.Finisher()(container)
}
