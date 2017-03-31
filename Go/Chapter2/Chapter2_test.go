package Chapter2

import (
	"testing"
	"reflect"
)

func TestRemoveDupes(t *testing.T) {
	list := newList("h","e","l","l","o")
	list.removeDupes()

	actual := list.toArray()
	expected  := []string{"h","e","l","o"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("RemoveDupes didn't work. Expected: %v Actual: %v \n", expected, actual)
	}

	list = newList("h","e","l","l","o")
	removeDupesRecursively(list.head, "l")
	actual = list.toArray()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("RemoveDupes didn't work. Expected: %v Actual: %v \n", expected, actual)
	}
}

func TestKthToLast(t *testing.T) {
	list := newList("h","e","l","2nd","o")
	expected := "2nd"
	actual := KthToLast(list, 2)

	if actual != expected {
		t.Errorf("KthToLast - Expected: %s Actual: %s \n", expected, actual)
	}

	list = newList("5th","e","l","2nd","o")
	expected = "5th"
	actual = KthToLast(list, 5)

	if actual != expected {
		t.Errorf("KthToLast - Expected: %s Actual: %s \n", expected, actual)
	}

	list = newList("5th","e","l","2nd","1st")
	actual = KthToLast(list, 1)
	expected = "1st"

	if actual != expected {
		t.Errorf("KthToLast - Expected: %s Actual: %s \n", expected, actual)
	}

}
