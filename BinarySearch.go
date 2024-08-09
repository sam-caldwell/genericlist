package genericlist

// BinarySearch - Perform sort (if needed) and binary search for a given value and set the current pointer to this node
func (list *LinkedList[T]) BinarySearch(item T) bool {
	if !list.sorted {
		list.Sort(func(a, b T) bool {
			return compareValues[T](a, b) == 1
		})
	}

	// Perform binary search
	left, right := 0, int(list.count)-1
	for left <= right {
		mid := left + (right-left)/2
		list.curr = list.head
		for i := 0; i < mid; i++ {
			if list.curr == nil {
				return false
			}
			list.curr = list.curr.nextPtr
		}
		if list.curr == nil {
			return false
		}

		if compareValues(list.curr.data, item) == 0 {
			return true
		} else if compareValues(list.curr.data, item) == -1 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
