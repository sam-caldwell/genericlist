package genericlist

const (
	equals      = 0
	greaterThan = 1
	lessThan    = -1
)

// Compare - Compare two generic values and return integer value result.
//
// output:
//
//	0  : a == b
//	+1 : a > b
//	-1 : a < b
func compareValues[T any](a, b T) int {
	var o T
	switch any(o).(type) {
	case int:
		if lhs, rhs := any(a).(int), any(b).(int); lhs == rhs {
			return equals
		} else if lhs > rhs {
			return greaterThan
		} else {
			return lessThan
		}
	case int8:
		if lhs, rhs := any(a).(int8), any(b).(int8); lhs == rhs {
			return equals
		} else if lhs > rhs {
			return greaterThan
		} else {
			return lessThan
		}
	case int16:
		if lhs, rhs := any(a).(int16), any(b).(int16); lhs == rhs {
			return equals
		} else if lhs > rhs {
			return greaterThan
		} else {
			return lessThan
		}
	case int32:
		if lhs, rhs := any(a).(int32), any(b).(int32); lhs == rhs {
			return equals
		} else if lhs > rhs {
			return greaterThan
		} else {
			return lessThan
		}
	case int64:
		if lhs, rhs := any(a).(int64), any(b).(int64); lhs == rhs {
			return equals
		} else if lhs > rhs {
			return greaterThan
		} else {
			return lessThan
		}

	case uint:
		if lhs, rhs := any(a).(uint), any(b).(uint); lhs == rhs {
			return equals
		} else if lhs > rhs {
			return greaterThan
		} else {
			return lessThan
		}
	case uint8:
		if lhs, rhs := any(a).(uint8), any(b).(uint8); lhs == rhs {
			return equals
		} else if lhs > rhs {
			return greaterThan
		} else {
			return lessThan
		}
	case uint16:
		if lhs, rhs := any(a).(uint16), any(b).(uint16); lhs == rhs {
			return equals
		} else if lhs > rhs {
			return greaterThan
		} else {
			return lessThan
		}
	case uint32:
		if lhs, rhs := any(a).(uint32), any(b).(uint32); lhs == rhs {
			return equals
		} else if lhs > rhs {
			return greaterThan
		} else {
			return lessThan
		}
	case uint64:
		if lhs, rhs := any(a).(uint64), any(b).(uint64); lhs == rhs {
			return equals
		} else if lhs > rhs {
			return greaterThan
		} else {
			return lessThan
		}
	case float32:
		if lhs, rhs := any(a).(float32), any(b).(float32); lhs == rhs {
			return equals
		} else if lhs > rhs {
			return greaterThan
		} else {
			return lessThan
		}
	case float64:
		if lhs, rhs := any(a).(float64), any(b).(float64); lhs == rhs {
			return equals
		} else if lhs > rhs {
			return greaterThan
		} else {
			return lessThan
		}
	case string:
		if lhs, rhs := any(a).(string), any(b).(string); lhs == rhs {
			return equals
		} else if lhs > rhs {
			return greaterThan
		} else {
			return lessThan
		}
	default:
		panic("BinarySearch only works for comparable types")
	}
	return 0
}
