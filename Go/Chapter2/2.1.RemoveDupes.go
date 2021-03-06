package Chapter2

// Remove Dups: Write code to remove duplicates from an unsorted linked list.
// FOLLOW UP
// How would you solve this problem if a temporary bu er is not allowed?
func (list *LinkedList) removeDupes() {
	set := make(map[int]bool)
	set[list.head.value] = true
	for node := list.head; node.next != nil; node = node.next {
		if set[node.next.value] {
			// delete the next node
			node.next = node.next.next
		} else {
			// hello
			set[node.next.value] = true
		}
	}
}
