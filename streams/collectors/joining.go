package collectors

import (
	"strings"

	"github.com/t0mj3dus0r/coffee"
)

type AnyString interface {
	~string
}

func Joining[T AnyString](delimiter, prefix, suffix string) coffee.Collector[T, *strings.Builder, string] {
	return &joining[T]{
		delimiter: delimiter,
		prefix:    prefix,
		suffix:    suffix,
		isFirst:   true,
	}
}

type joining[T AnyString] struct {
	delimiter, prefix, suffix string
	isFirst                   bool
}

// A function that folds a value into a mutable result container.
func (c *joining[T]) Accumulator() coffee.BiConsumer[*strings.Builder, T] {
	return func(b *strings.Builder, t T) {
		if !c.isFirst {
			b.WriteString(c.delimiter)
		} else {
			b.WriteString(c.prefix)
			c.isFirst = false
		}

		// fmt.Fprintf(b, "%v", t)
		b.WriteString(string(t))
	}
}

// A function that accepts two partial results and merges them.
func (c joining[T]) Combiner() func(a, b *strings.Builder) *strings.Builder {

	return func(a, b *strings.Builder) *strings.Builder {
		result := &strings.Builder{}

		result.WriteString(a.String())
		result.WriteString(b.String())

		return result
	}

}

// Perform the final transformation from the intermediate accumulation type A to the final result type R.
func (c joining[T]) Finisher() func(b *strings.Builder) string {
	return func(b *strings.Builder) string {
		b.WriteString(c.suffix)
		return b.String()
	}
}

// A function that creates and returns a new mutable result container.
func (c joining[T]) Supplier() coffee.Supplier[*strings.Builder] {
	return func() *strings.Builder {
		return &strings.Builder{}
	}
}
