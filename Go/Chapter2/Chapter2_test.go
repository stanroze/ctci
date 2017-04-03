package Chapter2

import (
	"testing"
	"reflect"
)

func TestRemoveDupes(t *testing.T) {
	list := newList(1, 2, 3, 4, 4, 5, 6)
	list.removeDupes()

	actual := list.toArray()
	expected := []int{1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("RemoveDupes didn't work. Expected: %v Actual: %v \n", expected, actual)
	}
}

func TestKthToLast(t *testing.T) {
	list := newList(1,2,3,4,5,6,9,8)
	expected := 9
	actual := KthToLast(list, 2)

	if actual != expected {
		t.Errorf("KthToLast - Expected: %s Actual: %s \n", expected, actual)
	}

	list = newList(1,2,3,4,5)
	expected = 1
	actual = KthToLast(list, 5)

	if actual != expected {
		t.Errorf("KthToLast - Expected: %s Actual: %s \n", expected, actual)
	}

	list = newList(1,2,3,4,5,6,7,8,9)
	actual = KthToLast(list, 1)
	expected = 9

	if actual != expected {
		t.Errorf("KthToLast - Expected: %s Actual: %s \n", expected, actual)
	}

}

func TestRemove(t *testing.T) {
	list := newList(1,2,3,2,2)
	var node *Node
	for node = list.head; node.value != 3; node = node.next {
	}
	list.Remove(node)
	actual := list.toArray()
	expected := []int{1,2,2,2}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("TestRemove - Expected: %v Actual: %v \n", expected, actual)
	}
	list = newList(1,4,2,2,2)
	for node = list.head; node.value != 4; node = node.next {
	}
	list.Remove(node)
	actual = list.toArray()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("TestRemove - Expected: %v Actual: %v \n", expected, actual)
	}

}
