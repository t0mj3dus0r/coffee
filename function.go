package coffee

type Function[T, R] func(T) R

type BiFunction[T, U, R] func(T, U) R