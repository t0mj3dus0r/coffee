package coffee

type Iterable[T any] interface {
	ForEach(action Consumer[T])
	Iterator() Iterator[T]
}
