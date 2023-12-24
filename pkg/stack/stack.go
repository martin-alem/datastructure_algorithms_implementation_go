package stack

import "datastructures_algorithms/pkg/linked_list/single"

type Stack[T any] struct {
	list *single.SinglyLinkedList[T]
}

func New[T any]() *Stack[T] {
	return &Stack[T]{
		list: single.New[T](),
	}
}

func Initialize[T any](list []T) *Stack[T] {
	return &Stack[T]{
		list: single.Initialize[T](list, true),
	}
}

func Initializer[T any](elements ...T) *Stack[T] {
	return &Stack[T]{
		list: single.Initializer[T](true, elements...),
	}
}

func (s *Stack[T]) Push(data T) {
	s.list.Insert(data, true)
}

func (s *Stack[T]) Pop() (bool, error) {
	value, err := s.list.Remove(0)
	return value, err
}

func (s *Stack[T]) Peek() (T, error) {
	value, err := s.list.Get(0)
	return value, err
}

func (s *Stack[T]) IsEmpty() bool {
	return s.list.IsEmpty()
}
