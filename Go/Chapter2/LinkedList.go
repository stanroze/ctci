package Chapter2

// Simple Linked List that stores strings
type Node struct {
	next *Node
	value string
}

type LinkedList struct {
	head *Node
}

func newList(value string) *LinkedList {
	return &LinkedList{&Node{nil, value}}
}

func (list *LinkedList) append(value string)  {
	var node *Node
	for node = list.head; node.next != nil; node = node.next {}
	node.next = &Node{nil, value}
}

func (list *LinkedList) toArray() []string {
	if list == nil {
		return nil
	}
	arr := make([]string, 0,0)
	for node := list.head;  node != nil; node = node.next{
		arr = append(arr, node.value)
	}

	return arr
}
