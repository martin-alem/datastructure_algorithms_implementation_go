package test

import (
	"datastructures_algorithms/pkg/linked_list/single"
	"testing"
)

func TestNew(t *testing.T) {
	sl := single.New[int]()
	if sl.GetSize() != 0 {
		t.Errorf("Expected %d but got %d", 0, sl.GetSize())
	}
}

func TestInitializeWithNonEmptySlice(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6}
	sl := single.Initialize(list)
	if sl.GetSize() != len(list) {
		t.Errorf("Expected %d but got %d", len(list), sl.GetSize())
	}
}

func TestInitializeWithEmptySlice(t *testing.T) {
	var list []int
	sl := single.Initialize(list)
	if sl.GetSize() != len(list) {
		t.Errorf("Expected %d but got %d", len(list), sl.GetSize())
	}
}

func TestInitializerWithArgs(t *testing.T) {
	sl := single.Initializer(false, 4, 5, 3, 4, 5, 3)
	if sl.GetSize() != 6 {
		t.Errorf("Expected %d but got %d", 6, sl.GetSize())
	}
}

func TestInitializerWithNoArgs(t *testing.T) {
	sl := single.Initializer[int](false)
	if sl.GetSize() != 0 {
		t.Errorf("Expected %d but got %d", 0, sl.GetSize())
	}
}

func TestInsertWithFrontTrueWhenListEmpty(t *testing.T) {
	sl := single.New[int]()
	item := 8
	sl.Insert(item, true)
	if sl.GetSize() != 1 {
		t.Errorf("Expect %d but got %d", 1, sl.GetSize())
	}

	value, _ := sl.Get(0)

	if value != item {
		t.Errorf("Expected %d but got %d", item, value)
	}

}

func TestInsertWithFrontTrueWhenListNotEmpty(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6}
	sl := single.Initialize(list)
	item := 8
	sl.Insert(item, true)
	if sl.GetSize() != 7 {
		t.Errorf("Expect %d but got %d", 7, sl.GetSize())
	}

	value, _ := sl.Get(0)

	if value != item {
		t.Errorf("Expected %d but got %d", item, value)
	}

}

func TestInsertWithFrontFalseWhenListEmpty(t *testing.T) {
	sl := single.New[int]()
	item := 8
	sl.Insert(item, false)
	if sl.GetSize() != 1 {
		t.Errorf("Expect %d but got %d", 1, sl.GetSize())
	}

	value, _ := sl.Get(0)

	if value != item {
		t.Errorf("Expected %d but got %d", item, value)
	}

}

func TestInsertWithFrontFalseWhenListNotEmpty(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6}
	sl := single.Initialize(list)
	item := 8
	sl.Insert(item, false)
	if sl.GetSize() != 7 {
		t.Errorf("Expect %d but got %d", 7, sl.GetSize())
	}

	value, _ := sl.Get(6)

	if value != item {
		t.Errorf("Expected %d but got %d", item, value)
	}

}

func TestRemoveWithEmptyList(t *testing.T) {
	sl := single.New[int]()

	ok, _ := sl.Remove(4)

	if ok != false {
		t.Errorf("Expected %v but got %v", false, ok)
	}

}

func TestRemoveWithOneElement(t *testing.T) {
	list := []int{1}
	sl := single.Initialize(list)

	ok, _ := sl.Remove(0)
	_, _ = sl.Remove(0)
	_, _ = sl.Remove(0)

	if ok != true || sl.IsEmpty() != true {
		t.Errorf("Expected %v %v but got %v %v", true, true, ok, sl.IsEmpty())
	}
}

func TestRemoveWithInvalidPosition(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	sl := single.Initialize(list)

	ok1, _ := sl.Remove(5)
	ok2, _ := sl.Remove(-1)

	if ok1 != false && ok2 != false {
		t.Errorf("Expected %v %v but got %v %v", false, false, ok1, ok2)
	}
}

func TestRemoveAtBeginning(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	sl := single.Initialize(list)

	startOk, _ := sl.Remove(0)
	startValue, _ := sl.Get(0)
	if startOk != true || startValue != 2 {
		t.Errorf("Expected %v %d but got %v %d", true, 2, startOk, startValue)
	}
}

func TestRemoveAtEnd(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	sl := single.Initialize(list)

	endOk, _ := sl.Remove(sl.GetSize() - 1)
	endValue, _ := sl.Get(sl.GetSize() - 1)
	if endOk != true || endValue != 4 {
		t.Errorf("Expected %v %d but got %v %d", true, 4, endOk, endValue)
	}
}

func TestRemoveWithValidPosition(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	sl := single.Initialize(list)

	ok, _ := sl.Remove(2)
	value, _ := sl.Get(2)
	if ok != true || value != 4 {
		t.Errorf("Expected %v %d but got %v %d", true, 4, ok, value)
	}

}

func TestGetWithEmptyList(t *testing.T) {
	sl := single.New[int]()

	value, err := sl.Get(4)

	if value != 0 || err == nil {
		t.Errorf("Expected %d %v but got %d %v", 0, err.Error(), value, err.Error())
	}
}

func TestGetWithInvalidPosition(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	sl := single.Initialize(list)

	value, err := sl.Get(5)

	if value != 0 || err == nil {
		t.Errorf("Expected %d %v but got %d %v", 0, err.Error(), value, err.Error())
	}
}

func TestGetWithValidPosition(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	sl := single.Initialize(list)

	value, err := sl.Get(3)

	if value != 4 || err != nil {
		t.Errorf("Expected %d %v but got %d %v", 0, nil, value, err.Error())
	}
}

func TestGetSizeWithEmptyList(t *testing.T) {
	sl := single.New[int]()

	value := sl.GetSize()

	if value != 0 {
		t.Errorf("Expected %d but got %d", 0, value)
	}
}

func TestGetSize(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	sl := single.Initialize(list)

	value := sl.GetSize()

	if value != len(list) {
		t.Errorf("Expected %d but got %d", len(list), value)
	}
}

func TestIsEmptyWithEmptyList(t *testing.T) {
	sl := single.New[int]()

	value := sl.IsEmpty()

	if value != true {
		t.Errorf("Expected %v but got %v", true, value)
	}
}

func TestIsEmpty(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	sl := single.Initialize(list)

	value := sl.IsEmpty()

	if value != false {
		t.Errorf("Expected %v but got %v", false, value)
	}
}
