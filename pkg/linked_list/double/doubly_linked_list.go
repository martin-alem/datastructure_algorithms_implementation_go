package double

import "errors"

type node[T any] struct {
	data T
	next *node[T]
	prev *node[T]
}

type DoublyLinkedList[T any] struct {
	head *node[T]
	tail *node[T]
	size int
}

func New[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func Initialize[T any](list []T, front ...bool) *DoublyLinkedList[T] {
	if len(list) == 0 {
		return New[T]()
	}

	dl := New[T]()

	if len(front) == 0 {
		front = []bool{false}
	}

	for _, item := range list {
		dl.Insert(item, front[0])
	}

	return dl
}

func Initializer[T any](front bool, elements ...T) *DoublyLinkedList[T] {
	if len(elements) == 0 {
		return New[T]()
	}

	dl := Initialize[T](elements, front)

	return dl
}

func (dl *DoublyLinkedList[T]) Insert(data T, front bool) {
	node := new(node[T])
	node.data = data
	node.next = nil
	node.prev = nil

	if dl.head != nil {
		if front == true {
			node.next = dl.head
			dl.head.prev = node
			dl.head = node
		} else {
			node.prev = dl.tail
			dl.tail.next = node
			dl.tail = node
		}
	} else {
		dl.head = node
		dl.tail = node
	}

	dl.size += 1
}

func (dl *DoublyLinkedList[T]) Remove(index int) (bool, error) {
	if dl.head == nil {
		return false, errors.New("list is empty")
	}

	if index < 0 || index >= dl.size {
		return false, errors.New("out of range")
	}

	currentNode := dl.head
	var prevNode *node[T]

	if index == 0 {
		dl.head = currentNode.next
		if dl.head != nil {
			currentNode.next = nil
			dl.head.prev = nil
		}
		dl.size -= 1
		return true, nil
	}

	if index == dl.size-1 {
		currentNode = dl.tail
		dl.tail = currentNode.prev
		currentNode.prev = nil
		dl.size -= 1
		return true, nil
	}

	for i := 0; i < index; i++ {
		prevNode = currentNode
		currentNode = currentNode.next
	}

	prevNode.next = currentNode.next
	currentNode.next.prev = prevNode

	currentNode.next = nil
	currentNode.prev = nil
	dl.size -= 1
	return true, nil
}

func (dl *DoublyLinkedList[T]) Get(index int) (T, error) {
	if dl.head == nil {
		return *new(T), errors.New("list is empty")
	}

	if index < 0 || index >= dl.size {
		return *new(T), errors.New("out of range")
	}

	currentNode := dl.head

	for i := 0; i < index; i++ {
		currentNode = currentNode.next
	}

	return currentNode.data, nil
}

func (dl *DoublyLinkedList[T]) GetSize() int {
	return dl.size
}

func (dl *DoublyLinkedList[T]) IsEmpty() bool {
	return dl.size == 0
}
