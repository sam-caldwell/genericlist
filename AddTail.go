package genericlist

// AddTail - Insert a node at the end of the linked list.
func (list *LinkedList[T]) AddTail(data T) {

	list.lock.Lock()
	defer list.lock.Unlock()

	newNode := &LinkedListNode[T]{data: data}

	if list.tail == nil { // List is empty
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.prevPtr = list.tail
		list.tail.nextPtr = newNode
		list.tail = newNode
	}
	if list.curr == nil {
		list.curr = newNode
	}

	list.count++ // Increment count
	list.sorted = false
}
