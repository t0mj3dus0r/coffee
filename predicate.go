package coffee

type Predicate[T any] interface {
	// Returns a composed predicate that represents a short-circuiting logical AND of this predicate and another.
	And(Predicate[T]) Predicate[T]
	// Returns a predicate that represents the logical negation of this predicate.
	Negate() Predicate[T]
	// Returns a composed predicate that represents a short-circuiting logical OR of this predicate and another.
	Or(Predicate[T]) Predicate[T]
	// Evaluates this predicate on the given argument.
	Test(T) bool
}
