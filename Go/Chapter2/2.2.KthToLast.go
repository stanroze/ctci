package Chapter2

// Return Kth to Last: Implement an algorithm to find the kth to last element
// of a singly linked list.
func KthToLast(list *LinkedList, k int) int {
	node := list.head
	for i := 0; i < k; i++ {
		if node == nil {
			return -1
		}

		node = node.next
	}
	var kth *Node
	for kth = list.head; node != nil; node, kth = node.next, kth.next {
	}

	return kth.value
}
