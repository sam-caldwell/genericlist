package genericlist

import "testing"

func TestCompareValues(t *testing.T) {

	tests := []struct {
		name     string
		a, b     any
		expected int
	}{
		// Integer comparisons
		{"int: a == b", 5, 5, 0},
		{"int: a > b", 7, 5, 1},
		{"int: a < b", 3, 5, -1},

		// Unsigned integer comparisons
		{"uint: a == b", uint(5), uint(5), 0},
		{"uint: a > b", uint(7), uint(5), 1},
		{"uint: a < b", uint(3), uint(5), -1},

		// Float comparisons
		{"float: a == b", 5.0, 5.0, 0},
		{"float: a > b", 7.1, 5.2, 1},
		{"float: a < b", 3.3, 5.4, -1},

		// String comparisons
		{"string: a == b", "apple", "apple", 0},
		{"string: a > b", "banana", "apple", 1},
		{"string: a < b", "apple", "banana", -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got int
			switch a := tt.a.(type) {
			case int:
				got = compareValues(a, tt.b.(int))
			case uint:
				got = compareValues(a, tt.b.(uint))
			case float64:
				got = compareValues(a, tt.b.(float64))
			case string:
				got = compareValues(a, tt.b.(string))
			default:
				t.Fatalf("unexpected type %T", a)
			}

			if got != tt.expected {
				t.Errorf("compareValues(%v, %v) = %d; want %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}
