package test

import (
	"datastructures_algorithms/pkg/linked_list/circle"
	"testing"
)

func TestNew(t *testing.T) {
	cl := circle.New[int]()
	if cl.GetSize() != 0 {
		t.Errorf("Expected %d but got %d", 0, cl.GetSize())
	}
}

func TestInitializeWithNonEmptySlice(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6}
	cl := circle.Initialize(list)
	if cl.GetSize() != len(list) {
		t.Errorf("Expected %d but got %d", len(list), cl.GetSize())
	}
}

func TestInitializeWithEmptySlice(t *testing.T) {
	var list []int
	cl := circle.Initialize(list)
	if cl.GetSize() != len(list) {
		t.Errorf("Expected %d but got %d", len(list), cl.GetSize())
	}
}

func TestInitializerWithArgs(t *testing.T) {
	cl := circle.Initializer(4, 5, 3, 4, 5, 3)
	if cl.GetSize() != 6 {
		t.Errorf("Expected %d but got %d", 6, cl.GetSize())
	}
}

func TestInsertWithFrontTrueWhenListEmpty(t *testing.T) {
	cl := circle.New[int]()
	item := 8
	cl.Insert(item, true)
	if cl.GetSize() != 1 {
		t.Errorf("Expect %d but got %d", 1, cl.GetSize())
	}

	value, _ := cl.Get(0)

	if value != item {
		t.Errorf("Expected %d but got %d", item, value)
	}

}

func TestInsertWithFrontTrueWhenListNotEmpty(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6}
	cl := circle.Initialize(list)
	item := 8
	cl.Insert(item, true)
	if cl.GetSize() != 7 {
		t.Errorf("Expect %d but got %d", 7, cl.GetSize())
	}

	value, _ := cl.Get(0)

	if value != item {
		t.Errorf("Expected %d but got %d", item, value)
	}

}

func TestInsertWithFrontFalseWhenListEmpty(t *testing.T) {
	cl := circle.New[int]()
	item := 8
	cl.Insert(item, false)
	if cl.GetSize() != 1 {
		t.Errorf("Expect %d but got %d", 1, cl.GetSize())
	}

	value, _ := cl.Get(0)

	if value != item {
		t.Errorf("Expected %d but got %d", item, value)
	}

}

func TestInsertWithFrontFalseWhenListNotEmpty(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6}
	cl := circle.Initialize(list)
	item := 8
	cl.Insert(item, false)
	if cl.GetSize() != 7 {
		t.Errorf("Expect %d but got %d", 7, cl.GetSize())
	}

	value, _ := cl.Get(6)

	if value != item {
		t.Errorf("Expected %d but got %d", item, value)
	}

}

func TestRemoveWithEmptyList(t *testing.T) {
	cl := circle.New[int]()

	ok, _ := cl.Remove(4)

	if ok != false {
		t.Errorf("Expected %v but got %v", false, ok)
	}

}

func TestRemoveWithOneElement(t *testing.T) {
	list := []int{1}
	cl := circle.Initialize(list)

	ok, _ := cl.Remove(0)
	_, _ = cl.Remove(0)
	_, _ = cl.Remove(0)

	if ok != true || cl.IsEmpty() != true {
		t.Errorf("Expected %v %v but got %v %v", true, true, ok, cl.IsEmpty())
	}
}

func TestRemoveWithInvalidPosition(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	cl := circle.Initialize(list)

	ok1, _ := cl.Remove(5)
	ok2, _ := cl.Remove(-1)

	if ok1 != false && ok2 != false {
		t.Errorf("Expected %v %v but got %v %v", false, false, ok1, ok2)
	}
}

func TestRemoveAtBeginning(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	cl := circle.Initialize(list)

	startOk, _ := cl.Remove(0)
	startValue, _ := cl.Get(0)
	if startOk != true || startValue != 2 {
		t.Errorf("Expected %v %d but got %v %d", true, 2, startOk, startValue)
	}
}

func TestRemoveAtEnd(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	cl := circle.Initialize(list)

	endOk, _ := cl.Remove(cl.GetSize() - 1)
	endValue, _ := cl.Get(cl.GetSize() - 1)
	if endOk != true || endValue != 4 {
		t.Errorf("Expected %v %d but got %v %d", true, 4, endOk, endValue)
	}
}

func TestRemoveWithValidPosition(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	cl := circle.Initialize(list)

	ok, _ := cl.Remove(2)
	value, _ := cl.Get(2)
	if ok != true || value != 4 {
		t.Errorf("Expected %v %d but got %v %d", true, 4, ok, value)
	}

}

func TestGetWithEmptyList(t *testing.T) {
	cl := circle.New[int]()

	value, err := cl.Get(4)

	if value != 0 || err == nil {
		t.Errorf("Expected %d %v but got %d %v", 0, err.Error(), value, err.Error())
	}
}

func TestGetWithInvalidPosition(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	cl := circle.Initialize(list)

	value, err := cl.Get(5)

	if value != 0 || err == nil {
		t.Errorf("Expected %d %v but got %d %v", 0, err.Error(), value, err.Error())
	}
}

func TestGetWithValidPosition(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	cl := circle.Initialize(list)

	value, err := cl.Get(3)

	if value != 4 || err != nil {
		t.Errorf("Expected %d %v but got %d %v", 0, nil, value, err.Error())
	}
}

func TestGetSizeWithEmptyList(t *testing.T) {
	cl := circle.New[int]()

	value := cl.GetSize()

	if value != 0 {
		t.Errorf("Expected %d but got %d", 0, value)
	}
}

func TestGetSize(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	cl := circle.Initialize(list)

	value := cl.GetSize()

	if value != len(list) {
		t.Errorf("Expected %d but got %d", len(list), value)
	}
}

func TestIsEmptyWithEmptyList(t *testing.T) {
	cl := circle.New[int]()

	value := cl.IsEmpty()

	if value != true {
		t.Errorf("Expected %v but got %v", true, value)
	}
}

func TestIsEmpty(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	cl := circle.Initialize(list)

	value := cl.IsEmpty()

	if value != false {
		t.Errorf("Expected %v but got %v", false, value)
	}
}
