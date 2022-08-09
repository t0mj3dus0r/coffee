package set

import (
	"github.com/t0mj3dus0r/coffee"
	"github.com/t0mj3dus0r/coffee/list"
)

func Of[T comparable](t ...T) coffee.Set[T] {
	instance := NewHashSet[T]()

	instance.AddAll(list.FromArray(t))

	return instance
}
