package Chapter2

import (
	"testing"

	"reflect"
)

func TestRemoveDupes(t *testing.T) {
	list := newList("h")
	list.append("e")
	list.append("l")
	list.append("l")
	list.append("o")
	list.removeDupes()

	actual := list.toArray()
	expected  := []string{"h","e","l","o"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("RemoveDupes didn't work. Expected: %v Actual: %v \n", expected, actual)
	}

	list = newList("h")
	list.append("e")
	list.append("l")
	list.append("l")
	list.append("o")
	removeDupesRecursively(list.head, "l")
	actual = list.toArray()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("RemoveDupes didn't work. Expected: %v Actual: %v \n", expected, actual)
	}


}
