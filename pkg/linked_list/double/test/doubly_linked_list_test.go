package test

import (
	"datastructures_algorithms/pkg/linked_list/double"
	"testing"
)

func TestNew(t *testing.T) {
	dl := double.New[int]()
	if dl.GetSize() != 0 {
		t.Errorf("Expected %d but got %d", 0, dl.GetSize())
	}
}

func TestInitializeWithNonEmptySlice(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6}
	dl := double.Initialize(list)
	if dl.GetSize() != len(list) {
		t.Errorf("Expected %d but got %d", len(list), dl.GetSize())
	}
}

func TestInitializeWithEmptySlice(t *testing.T) {
	var list []int
	dl := double.Initialize(list)
	if dl.GetSize() != len(list) {
		t.Errorf("Expected %d but got %d", len(list), dl.GetSize())
	}
}

func TestInitializerWithArgs(t *testing.T) {
	dl := double.Initializer(4, 5, 3, 4, 5, 3)
	if dl.GetSize() != 6 {
		t.Errorf("Expected %d but got %d", 6, dl.GetSize())
	}
}

func TestInsertWithFrontTrueWhenListEmpty(t *testing.T) {
	dl := double.New[int]()
	item := 8
	dl.Insert(item, true)
	if dl.GetSize() != 1 {
		t.Errorf("Expect %d but got %d", 1, dl.GetSize())
	}

	value, _ := dl.Get(0)

	if value != item {
		t.Errorf("Expected %d but got %d", item, value)
	}

}

func TestInsertWithFrontTrueWhenListNotEmpty(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6}
	dl := double.Initialize(list)
	item := 8
	dl.Insert(item, true)
	if dl.GetSize() != 7 {
		t.Errorf("Expect %d but got %d", 7, dl.GetSize())
	}

	value, _ := dl.Get(0)

	if value != item {
		t.Errorf("Expected %d but got %d", item, value)
	}

}

func TestInsertWithFrontFalseWhenListEmpty(t *testing.T) {
	dl := double.New[int]()
	item := 8
	dl.Insert(item, false)
	if dl.GetSize() != 1 {
		t.Errorf("Expect %d but got %d", 1, dl.GetSize())
	}

	value, _ := dl.Get(0)

	if value != item {
		t.Errorf("Expected %d but got %d", item, value)
	}

}

func TestInsertWithFrontFalseWhenListNotEmpty(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6}
	dl := double.Initialize(list)
	item := 8
	dl.Insert(item, false)
	if dl.GetSize() != 7 {
		t.Errorf("Expect %d but got %d", 7, dl.GetSize())
	}

	value, _ := dl.Get(6)

	if value != item {
		t.Errorf("Expected %d but got %d", item, value)
	}

}

func TestRemoveWithEmptyList(t *testing.T) {
	dl := double.New[int]()

	ok, _ := dl.Remove(4)

	if ok != false {
		t.Errorf("Expected %v but got %v", false, ok)
	}

}

func TestRemoveWithOneElement(t *testing.T) {
	list := []int{1}
	dl := double.Initialize(list)

	ok, _ := dl.Remove(0)
	_, _ = dl.Remove(0)
	_, _ = dl.Remove(0)

	if ok != true || dl.IsEmpty() != true {
		t.Errorf("Expected %v %v but got %v %v", true, true, ok, dl.IsEmpty())
	}
}

func TestRemoveWithInvalidPosition(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	dl := double.Initialize(list)

	ok1, _ := dl.Remove(5)
	ok2, _ := dl.Remove(-1)

	if ok1 != false && ok2 != false {
		t.Errorf("Expected %v %v but got %v %v", false, false, ok1, ok2)
	}
}

func TestRemoveAtBeginning(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	dl := double.Initialize(list)

	startOk, _ := dl.Remove(0)
	startValue, _ := dl.Get(0)
	if startOk != true || startValue != 2 {
		t.Errorf("Expected %v %d but got %v %d", true, 2, startOk, startValue)
	}
}

func TestRemoveAtEnd(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	dl := double.Initialize(list)

	endOk, _ := dl.Remove(dl.GetSize() - 1)
	endValue, _ := dl.Get(dl.GetSize() - 1)
	if endOk != true || endValue != 4 {
		t.Errorf("Expected %v %d but got %v %d", true, 4, endOk, endValue)
	}
}

func TestRemoveWithValidPosition(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	dl := double.Initialize(list)

	ok, _ := dl.Remove(2)
	value, _ := dl.Get(2)
	if ok != true || value != 4 {
		t.Errorf("Expected %v %d but got %v %d", true, 4, ok, value)
	}

}

func TestGetWithEmptyList(t *testing.T) {
	dl := double.New[int]()

	value, err := dl.Get(4)

	if value != 0 || err == nil {
		t.Errorf("Expected %d %v but got %d %v", 0, err.Error(), value, err.Error())
	}
}

func TestGetWithInvalidPosition(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	dl := double.Initialize(list)

	value, err := dl.Get(5)

	if value != 0 || err == nil {
		t.Errorf("Expected %d %v but got %d %v", 0, err.Error(), value, err.Error())
	}
}

func TestGetWithValidPosition(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	dl := double.Initialize(list)

	value, err := dl.Get(3)

	if value != 4 || err != nil {
		t.Errorf("Expected %d %v but got %d %v", 0, nil, value, err.Error())
	}
}

func TestGetSizeWithEmptyList(t *testing.T) {
	dl := double.New[int]()

	value := dl.GetSize()

	if value != 0 {
		t.Errorf("Expected %d but got %d", 0, value)
	}
}

func TestGetSize(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	dl := double.Initialize(list)

	value := dl.GetSize()

	if value != len(list) {
		t.Errorf("Expected %d but got %d", len(list), value)
	}
}

func TestIsEmptyWithEmptyList(t *testing.T) {
	dl := double.New[int]()

	value := dl.IsEmpty()

	if value != true {
		t.Errorf("Expected %v but got %v", true, value)
	}
}

func TestIsEmpty(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	dl := double.Initialize(list)

	value := dl.IsEmpty()

	if value != false {
		t.Errorf("Expected %v but got %v", false, value)
	}
}
