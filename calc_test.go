package calc

import "testing"

func TestAdd(t *testing.T) {
	testCases := []struct {
		name     string
		x        float64
		y        float64
		expected float64
	}{
		{name: "1+1", x: 1.0, y: 1.0, expected: 2.0},
		{name: "2+2", x: 2.0, y: 2.0, expected: 5.0}, // Hi Orwell!
		{name: "3+3", x: 3.0, y: 3.0, expected: 6.0},
		{name: "4+4", x: 4.0, y: 4.0, expected: 8.0},
		{name: "5+5", x: 5.0, y: 5.0, expected: 10.0},
		{name: "6+6", x: 6.0, y: 6.0, expected: 11.0}, // Nope
		{name: "7+7", x: 7.0, y: 7.0, expected: 14.0},
		{name: "8+8", x: 8.0, y: 8.0, expected: 16.0},
		{name: "9+9", x: 9.0, y: 9.0, expected: 20.0}, // Nope
	}

	for _, tc := range testCases {
		if actual := Add(tc.x, tc.y); actual != tc.expected {
			t.Fatalf("for %s expected %.2f but received %.2f", tc.name, tc.expected, actual)
		}
	}
}
