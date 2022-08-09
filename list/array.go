package list

import (
	"github.com/t0mj3dus0r/coffee"
	"github.com/t0mj3dus0r/coffee/iterator"
	"github.com/t0mj3dus0r/coffee/predicate"
	"github.com/t0mj3dus0r/coffee/streams"
	"golang.org/x/exp/slices"
)

func FromArray[T comparable](data []T) *ArrayList[T] {
	return &ArrayList[T]{data}
}

func NewArrayList[T comparable]() *ArrayList[T] {
	return &ArrayList[T]{data: make([]T, 0)}
}

type ArrayList[E comparable] struct {
	data []E
}

func (a *ArrayList[E]) Add(e E) bool {
	a.data = append(a.data, e)
	return true
}

func (a ArrayList[E]) AddAll(c coffee.Collection[E]) bool {
	c.ForEach(func(e E) {
		a.data = append(a.data, e)
	})
	return true
}

func (a *ArrayList[E]) Clear() {
	a.data = []E{}
}

func (a ArrayList[E]) Contains(e E) bool {
	return slices.Contains(a.data, e)
}

func (a ArrayList[E]) ContainsAll(c coffee.Collection[E]) bool {
	// TODO rework the complexity

	iterator := c.Iterator()

	for iterator.HasNext() {
		if !a.Contains(iterator.Next()) {
			return false
		}
	}

	return true
}

func (a ArrayList[E]) IsEmpty() bool {
	return len(a.data) == 0
}

func (a *ArrayList[E]) Iterator() coffee.Iterator[E] {
	return iterator.List[E](a)
}

func (a *ArrayList[E]) Remove(e E) bool {

	var index int

	for _, v := range a.data {
		if v == e {
			a.data = slices.Delete(a.data, index, index)
			return true
		}

		index++
	}

	return false
}

func (a *ArrayList[E]) RemoveIf(filter coffee.Predicate[E]) bool {

	var modified bool
	var index int

	for _, v := range a.data {
		if filter.Test(v) {
			a.data = slices.Delete(a.data, index, index)
			modified = true
		}

		index++
	}

	return modified
}

func (a ArrayList[E]) RemoveAll(c coffee.Collection[E]) bool {
	// TODO rework the complexity

	var hasModified bool

	a.RemoveIf(predicate.From(func(e E) bool {
		toRemove := c.Contains(e)
		if toRemove {
			hasModified = true
		}

		return toRemove
	}))

	return hasModified
}

func (a ArrayList[E]) Size() int {
	return len(a.data)
}

func (a ArrayList[E]) Stream() coffee.Stream[E] {
	return streams.FromArray(a.ToArray())
}

// TODO make it deep cloning
// WARNING this function creates Shallow clones
func (a ArrayList[E]) ToArray() []E {
	return slices.Clone(a.data)
}

func (a ArrayList[E]) ForEach(action coffee.Consumer[E]) {
	for _, v := range a.data {
		action(v)
	}
}

func (a ArrayList[E]) IndexOf(E) int {
	panic("not implemented yet")
}

func (a ArrayList[E]) RemoveByIndex(index int) E {
	panic("not implemented yet")
}

func (a ArrayList[E]) Set(index int, element E) E {
	panic("not implemented yet")
}

func (a ArrayList[E]) Get(index int) E {
	return a.data[index]
}
