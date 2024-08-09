package genericlist

// Sort - Sort the linked list in ascending or descending order.
func (list *LinkedList[T]) Sort(swap func(a, b T) bool) {
	list.lock.Lock()
	defer list.lock.Unlock()
	if list.count < 2 {
		return // No need to sort if the list has less than 2 elements
	}
	swapped := true
	for swapped {
		swapped = false
		current := list.head

		for current != nil && current.nextPtr != nil {
			// Compare based on the 'asc' parameter
			if swap(current.data, current.nextPtr.data) {
				current.data, current.nextPtr.data = current.nextPtr.data, current.data
				swapped = true
			}
			current = current.nextPtr
		}
	}
	list.sorted = true
}
