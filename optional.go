package coffee

type Optional[T any] interface {
	// If a value is present, and the value matches the given predicate, returns an Optional describing the value, otherwise returns an empty Optional.
	Filter(predicate Predicate[T]) Optional[T]
	// If a value is present, performs the given action with the value, otherwise does nothing.
	IfPresent(action Consumer[T])
	// If a value is present, performs the given action with the value, otherwise performs the given empty-based action.
	IfPresentOrElse(action Consumer[T], emptyAction func())
	// If a value is not present, returns true, otherwise false.
	IsEmpty() bool
	// If a value is present, returns true, otherwise false.
	IsPresent() bool
	// If a value is present, returns an Optional describing the value, otherwise returns an Optional produced by the supplying function.
	Or(supplier Supplier[Optional[T]]) Optional[T]
	// If a value is present, returns the value, otherwise returns other.
	OrElse(other T) T
	// If a value is present, returns the value, otherwise returns the result produced by the supplying function.
	OrElseGet(supplier Supplier[T]) T
	// If a value is present, returns the value, otherwise panics.
	OrElsePanic() T
}
