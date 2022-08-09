package iterator

import "github.com/t0mj3dus0r/coffee"

func List[E comparable](list coffee.List[E]) coffee.Iterator[E] {
	return &listIterator[E]{
		arrayList: list,
		index:     -1,
	}
}

type listIterator[E comparable] struct {
	arrayList coffee.List[E]

	index int
}

func (a *listIterator[E]) ForEachRemaining(action coffee.Consumer[E]) {
	for a.HasNext() {
		action(a.Next())
	}
}

func (a *listIterator[E]) HasNext() bool {
	return a.arrayList.Size()-1 > a.index

}

func (a *listIterator[E]) Next() E {
	if !a.HasNext() {
		var e E
		return e
	}

	a.index++

	return a.arrayList.Get(a.index)
}
