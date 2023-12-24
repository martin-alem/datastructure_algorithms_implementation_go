package single

import "errors"

type node[T any] struct {
	data T
	next *node[T]
}

type SinglyLinkedList[T any] struct {
	head *node[T]
	tail *node[T]
	size int
}

func New[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func Initialize[T any](list []T, front ...bool) *SinglyLinkedList[T] {
	if len(list) == 0 {
		return New[T]()
	}

	if len(front) == 0 {
		front = []bool{false}
	}
	sl := New[T]()

	for _, item := range list {
		sl.Insert(item, front[0])
	}

	return sl
}

func Initializer[T any](front bool, elements ...T) *SinglyLinkedList[T] {
	if len(elements) == 0 {
		return New[T]()
	}

	sl := Initialize[T](elements, front)

	return sl
}

func (sl *SinglyLinkedList[T]) Insert(data T, front bool) {
	node := new(node[T])
	node.data = data
	node.next = nil

	if sl.head != nil {
		if front == true {
			node.next = sl.head
			sl.head = node
		} else {
			sl.tail.next = node
			sl.tail = node
		}
	} else {
		sl.head = node
		sl.tail = node
	}

	sl.size += 1
}

func (sl *SinglyLinkedList[T]) Remove(index int) (bool, error) {
	if sl.head == nil {
		return false, errors.New("list is empty")
	}

	if index < 0 || index >= sl.size {
		return false, errors.New("out of range")
	}

	currentNode := sl.head
	var prevNode *node[T]

	if index == 0 {
		sl.head = currentNode.next
		if sl.head != nil {
			currentNode.next = nil
		}
		sl.size -= 1
		return true, nil
	}

	for i := 0; i < index; i++ {
		prevNode = currentNode
		currentNode = currentNode.next
	}

	prevNode.next = currentNode.next

	//Resetting the tail node
	if currentNode.next == nil {
		sl.tail = prevNode
	}
	currentNode.next = nil
	sl.size -= 1
	return true, nil
}

func (sl *SinglyLinkedList[T]) Get(index int) (T, error) {
	if sl.head == nil {
		return *new(T), errors.New("list is empty")
	}

	if index < 0 || index >= sl.size {
		return *new(T), errors.New("out of range")
	}

	currentNode := sl.head

	for i := 0; i < index; i++ {
		currentNode = currentNode.next
	}

	return currentNode.data, nil
}

func (sl *SinglyLinkedList[T]) GetSize() int {
	return sl.size
}

func (sl *SinglyLinkedList[T]) IsEmpty() bool {
	return sl.size == 0
}
