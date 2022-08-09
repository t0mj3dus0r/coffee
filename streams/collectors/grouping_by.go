package collectors

import (
	"github.com/t0mj3dus0r/coffee"
)

func GroupingBy[T, K comparable](classifier func(T) K) coffee.Collector[T, map[K][]T, map[K][]T] {
	return groupingBy[T, K]{classifier: classifier}
}

type groupingBy[T any, K comparable] struct {
	classifier func(T) K
}

// A function that folds a value into a mutable result container.
func (c groupingBy[T, K]) Accumulator() coffee.BiConsumer[map[K][]T, T] {
	return func(a map[K][]T, t T) {

		k := c.classifier(t)

		_, exists := a[k]
		if !exists {
			a[k] = make([]T, 0)
		}

		a[k] = append(a[k], t)
	}
}

// A function that accepts two partial results and merges them.
func (c groupingBy[T, K]) Combiner() func(a, b map[K][]T) map[K][]T {
	panic("not implemented") // TODO: Implement
}

// Perform the final transformation from the intermediate accumulation type A to the final result type R.
func (c groupingBy[T, K]) Finisher() func(a map[K][]T) map[K][]T {
	return func(a map[K][]T) map[K][]T {
		return a
	}
}

// A function that creates and returns a new mutable result container.
func (c groupingBy[T, K]) Supplier() coffee.Supplier[map[K][]T] {
	return func() map[K][]T {
		return make(map[K][]T)
	}
}
