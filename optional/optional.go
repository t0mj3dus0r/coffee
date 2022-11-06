package optional

import (
	"github.com/t0mj3dus0r/coffee"
)

func Of[T any](t T) coffee.Optional[T] {
	return optionalImpl[T]{
		data: t,
	}
}

func OfNullable[T any](t *T) coffee.Optional[T] {
	if t == nil {
		return optionalImpl[T]{
			isEmpty: true,
		}
	}

	return optionalImpl[T]{
		data: *t,
	}
}

func Empty[T any]() coffee.Optional[T] {
	return optionalImpl[T]{
		isEmpty: true,
	}
}

type optionalImpl[T any] struct {
	isEmpty bool
	data    T
}

// If a value is present, and the value matches the given predicate, returns an Optional describing the value, otherwise returns an empty Optional.
func (o optionalImpl[T]) Filter(predicate coffee.Predicate[T]) coffee.Optional[T] {
	if o.IsPresent() {
		if predicate.Test(o.data) {
			return o
		}
	}

	return Empty[T]()
}

// If a value is present, performs the given action with the value, otherwise does nothing.
func (o optionalImpl[T]) IfPresent(action coffee.Consumer[T]) {
	if o.IsPresent() {
		action(o.data)
	}
}

// If a value is present, performs the given action with the value, otherwise performs the given empty-based action.
func (o optionalImpl[T]) IfPresentOrElse(action coffee.Consumer[T], emptyAction func()) {
	if o.IsPresent() {
		action(o.data)
	} else {
		emptyAction()
	}
}

// If a value is not present, returns true, otherwise false.
func (o optionalImpl[T]) IsEmpty() bool {
	return o.isEmpty
}

// If a value is present, returns true, otherwise false.
func (o optionalImpl[T]) IsPresent() bool {
	return !o.isEmpty
}

// If a value is present, returns an Optional describing the value, otherwise returns an Optional produced by the supplying function.
func (o optionalImpl[T]) Or(supplier coffee.Supplier[coffee.Optional[T]]) coffee.Optional[T] {
	if o.IsPresent() {
		return o
	}

	return supplier()
}

// If a value is present, returns the value, otherwise returns other.
func (o optionalImpl[T]) OrElse(other T) T {
	if o.IsPresent() {
		return o.data
	}

	return other
}

// If a value is present, returns the value, otherwise returns the result produced by the supplying function.
func (o optionalImpl[T]) OrElseGet(supplier coffee.Supplier[T]) T {
	if o.IsPresent() {
		return o.data
	}

	return supplier()
}

// If a value is present, returns the value, otherwise panics.
func (o optionalImpl[T]) OrElsePanic() T {
	if o.IsPresent() {
		return o.data
	}

	panic("no value present in optional")
}
