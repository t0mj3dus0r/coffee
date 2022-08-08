package predicate

import "github.com/t0mj3dus0r/coffee"

func From[T any](predicateFunc func(T) bool) coffee.Predicate[T] {
	return predicate[T]{predicateFunc: predicateFunc}
}

type predicate[T any] struct {
	predicateFunc func(T) bool
}

// Returns a composed predicate that represents a short-circuiting logical AND of this predicate and another.
func (p predicate[T]) And(other coffee.Predicate[T]) coffee.Predicate[T] {
	return predicate[T]{
		predicateFunc: func(t T) bool {
			return p.Test(t) && other.Test(t)
		},
	}
}

// Returns a predicate that represents the logical negation of this predicate.
func (p predicate[T]) Negate() coffee.Predicate[T] {
	return predicate[T]{
		predicateFunc: func(t T) bool {
			return !p.Test(t)
		},
	}
}

// Returns a composed predicate that represents a short-circuiting logical OR of this predicate and another.
func (p predicate[T]) Or(other coffee.Predicate[T]) coffee.Predicate[T] {
	return predicate[T]{
		predicateFunc: func(t T) bool {
			return p.Test(t) || other.Test(t)
		},
	}
}

// Evaluates this predicate on the given argument.
func (p predicate[T]) Test(t T) bool {
	return p.predicateFunc(t)
}
