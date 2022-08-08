package set

import "github.com/t0mj3dus0r/coffee"

func Of[T comparable](t ...T) coffee.Set[T] {
	instance := NewHashSet[T]()

	instance.AddAll(t)

	return instance
}
