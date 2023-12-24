package test

import (
	"datastructures_algorithms/pkg/stack"
	"testing"
)

func TestNew(t *testing.T) {
	s := stack.New[int]()
	if s.IsEmpty() != true {
		t.Errorf("Expected %v but got %v", true, s.IsEmpty())
	}
}

func TestInitializeWithList(t *testing.T) {
	list := []int{2, 3, 4, 65}
	s := stack.Initialize(list)
	if s.IsEmpty() != false {
		t.Errorf("Expected %v but got %v", false, s.IsEmpty())
	}
}

func TestInitializeWithEmptyList(t *testing.T) {
	var list []int
	s := stack.Initialize(list)
	if s.IsEmpty() != true {
		t.Errorf("Expected %v but got %v", true, s.IsEmpty())
	}
}

func TestInitializerWithArgs(t *testing.T) {
	s := stack.Initializer(3, 4, 5, 3, 5)
	if s.IsEmpty() != false {
		t.Errorf("Expected %v but got %v", false, s.IsEmpty())
	}
}

func TestInitializerWithNoArgs(t *testing.T) {
	s := stack.Initializer[int]()
	if s.IsEmpty() != true {
		t.Errorf("Expected %v but got %v", true, s.IsEmpty())
	}
}

func TestPushWithEmptyStack(t *testing.T) {
	s := stack.New[int]()
	s.Push(5)
	value, err := s.Peek()
	if value != 5 || err != nil {
		t.Errorf("Expected %v %v but got %v %v", 5, nil, value, err)
	}
}

func TestPushWithElementsInStack(t *testing.T) {
	s := stack.Initializer(4, 5, 6, 2)
	s.Push(10)
	value, err := s.Peek()
	if value != 10 || err != nil {
		t.Errorf("Expected %v %v but got %v %v", 10, nil, value, err)
	}
}

func TestPopWithEmptyStack(t *testing.T) {
	s := stack.New[int]()
	ok, _ := s.Pop()
	if ok != false {
		t.Errorf("Expected %v but got %v", false, ok)
	}
}

func TestPopWithElementsInStack(t *testing.T) {
	s := stack.Initializer[int](3, 4, 6, 7, 8)
	ok, _ := s.Pop()
	if ok != true {
		t.Errorf("Expected %v but got %v", true, ok)
	}

	value, _ := s.Peek()
	if value != 7 {
		t.Errorf("Expected %v but got %v", 7, value)
	}
}

func TestPopWithOneElementInStack(t *testing.T) {
	s := stack.Initializer[int](3)
	ok, _ := s.Pop()
	if ok != true {
		t.Errorf("Expected %v but got %v", true, ok)
	}

	if s.IsEmpty() != true {
		t.Errorf("Expected %v but got %v", true, s.IsEmpty())
	}
}

func TestPeekWithEmptyStack(t *testing.T) {
	s := stack.New[int]()
	value, _ := s.Peek()
	if value != 0 {
		t.Errorf("Expected %v but got %v", 0, value)
	}
}

func TestPeekWithElementsInEmptyStack(t *testing.T) {
	s := stack.Initializer[int](3, 4, 6, 7, 8)
	value, _ := s.Peek()
	if value != 8 {
		t.Errorf("Expected %v but got %v", 8, value)
	}
}

func TestPeekWithOneElementInEmptyStack(t *testing.T) {
	s := stack.Initializer[int](8)
	value, _ := s.Peek()
	if value != 8 {
		t.Errorf("Expected %v but got %v", 8, value)
	}
}

func TestIsEmptyWithEmptyStack(t *testing.T) {
	s := stack.New[int]()
	if s.IsEmpty() != true {
		t.Errorf("Expected %v but got %v", true, s.IsEmpty())
	}
}

func TestIsEmptyWithElementsInStack(t *testing.T) {
	s := stack.Initializer[int](3, 4, 6, 7, 8)
	if s.IsEmpty() != false {
		t.Errorf("Expected %v but got %v", false, s.IsEmpty())
	}
}
