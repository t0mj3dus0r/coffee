package coffee

type Function[T, R any] func(T) R

type BiFunction[T, U, R any] func(T, U) R
