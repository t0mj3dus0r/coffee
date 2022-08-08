package coffee

type Consumer[T any] func(T)

type BiConsumer[T, U any] func(T, U)
