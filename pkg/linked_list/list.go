package linked_list

type List[T any] interface {
	Insert(data T, front bool)
	Remove(position int) (bool, error)
	Get(position int) (T, error)
	GetSize() int
	IsEmpty() bool
}
