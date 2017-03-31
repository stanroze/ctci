package Chapter2

// Delete Middle Node: Implement an algorithm to delete a node in the middle
// (i.e., any node but the  first and last node, not necessarily the exact middle)
// of a singly linked list, given only access to that node.
// EXAMPLE
// Input:the node c from the linked list a->b->c->d->e->f
// Result: nothing is returned, but the new linked list looks like a->b->d->e->f
func (list *LinkedList) Remove(node *Node){
	var item *Node
	for item = node; item.next.next != nil; item = item.next {
		item.value = item.next.value
	}
	item.value = item.next.value
	item.next = nil
}

