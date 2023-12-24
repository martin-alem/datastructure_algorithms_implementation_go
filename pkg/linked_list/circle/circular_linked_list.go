package circle

import "errors"

type node[T any] struct {
	data T
	next *node[T]
}

type CircularLinkedList[T any] struct {
	head *node[T]
	tail *node[T]
	size int
}

func New[T any]() *CircularLinkedList[T] {
	return &CircularLinkedList[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func Initialize[T any](list []T, front ...bool) *CircularLinkedList[T] {
	if len(list) == 0 {
		return New[T]()
	}

	if len(front) == 0 {
		front = []bool{false}
	}

	cl := New[T]()

	for _, v := range list {
		cl.Insert(v, front[0])
	}

	return cl
}

func Initializer[T any](front bool, elements ...T) *CircularLinkedList[T] {
	if len(elements) == 0 {
		return New[T]()
	}

	cl := Initialize[T](elements, front)

	return cl
}

func (cl *CircularLinkedList[T]) Insert(data T, front bool) {
	node := new(node[T])
	node.data = data
	node.next = nil

	if cl.head != nil {
		if front == false {
			cl.tail.next = node
			node.next = cl.head
			cl.tail = node
		} else {
			node.next = cl.head
			cl.head = node
		}
	} else {
		cl.head = node
		cl.tail = node
		node.next = cl.head
	}

	cl.size += 1
}

func (cl *CircularLinkedList[T]) Remove(index int) (bool, error) {
	if cl.head == nil {
		return false, errors.New("list is empty")
	}

	if index < 0 || index >= cl.size {
		return false, errors.New("out of range")
	}

	currentNode := cl.head
	var prevNode *node[T]

	if index == 0 {
		//Last node in list
		if currentNode.next == cl.head {
			cl.head = nil
			cl.tail = nil
		} else { // At least two nodes in the list
			cl.head = currentNode.next
			currentNode.next = nil
		}
		cl.size -= 1
		return true, nil
	}

	for i := 0; i < index; i++ {
		prevNode = currentNode
		currentNode = currentNode.next
	}

	//Removing last node
	if currentNode == cl.tail {
		prevNode.next = cl.head
		cl.tail = prevNode
	} else {
		//Removing nodes between first and last
		prevNode.next = currentNode.next
	}

	currentNode.next = nil
	cl.size -= 1
	return true, nil
}

func (cl *CircularLinkedList[T]) Get(index int) (T, error) {
	if cl.head == nil {
		return *new(T), errors.New("list is empty")
	}

	if index < 0 || index >= cl.size {
		return *new(T), errors.New("out of range")
	}

	currentNode := cl.head

	for i := 0; i < index; i++ {
		currentNode = currentNode.next
	}

	return currentNode.data, nil
}

func (cl *CircularLinkedList[T]) GetSize() int {
	return cl.size
}

func (cl *CircularLinkedList[T]) IsEmpty() bool {
	return cl.size == 0
}
