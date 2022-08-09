package coffee

type Collector[T, A, R any] interface {
	// A function that folds a value into a mutable result container.
	Accumulator() BiConsumer[A, T]
	// A function that accepts two partial results and merges them.
	Combiner() func(a, b A) A
	// Perform the final transformation from the intermediate accumulation type A to the final result type R.
	Finisher() func(A) R
	// A function that creates and returns a new mutable result container.
	Supplier() Supplier[A]
}
