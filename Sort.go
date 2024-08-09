package genericlist

// Sort - Sort the linked list in ascending or descending order.
func (list *LinkedList[T]) Sort(asc bool, less func(a, b T) bool) {
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
			if (asc && !less(current.data, current.nextPtr.data)) ||
				(!asc && less(current.data, current.nextPtr.data)) {
				// Swap the data
				current.data, current.nextPtr.data = current.nextPtr.data, current.data
				swapped = true
			}
			current = current.nextPtr
		}
	}
}
