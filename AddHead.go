package genericlist

// AddHead - Insert a node at the start of the linked list.
func (list *LinkedList[T]) AddHead(data T) {
	list.lock.Lock()
	defer list.lock.Unlock()

	newNode := &LinkedListNode[T]{data: data}

	if list.head == nil { // List is empty
		list.head = newNode
		list.tail = newNode
		list.curr = newNode
	} else {
		newNode.nextPtr = list.head
		list.head.prevPtr = newNode
		list.head = newNode
	}

	// Move curr to the new head
	if list.curr == nil {
		list.curr = newNode
	}

	list.count++ // Increment count
	list.sorted = false
}
