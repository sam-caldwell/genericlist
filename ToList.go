package genericlist

// ToList - Return a slice of node data from the linked list.
func (list *LinkedList[T]) ToList() []T {
	list.lock.Lock()
	defer list.lock.Unlock()
	var output []T
	list.curr = list.head
	for list.curr != nil {
		output = append(output, list.curr.data)
		if !list.moveNext() {
			break
		}
	}
	return output
}
