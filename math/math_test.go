package math

import (
	"testing"
)

type addTest struct {
	arg1, arg2, expected int
}

var addTests = []addTest{
	addTest{2, 3, 5},
	addTest{4, 8, 12},
	addTest{6, 9, 15},
	addTest{3, 10, 13},
}

var subtractTests = []addTest{
	addTest{10, 3, 7},
	addTest{9, 8, 1},
	addTest{6, 10, -4},
	addTest{1000, 1, 999},
}

func TestAdd(t *testing.T) {
	for _, test := range addTests {
		if output := Add(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

func TestSubtract(t *testing.T) {
	for _, test := range subtractTests {
		if output := Subtract(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
