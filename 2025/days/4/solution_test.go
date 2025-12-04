package main

import "testing"

func TestGetNeighbours(t *testing.T) {

	type inputs struct {
		x, y, maxX, maxY int
	}
	tests := []struct {
		testName string
		inputs
		expectedOutput [][]int
	}{
		{
			"top left",
			inputs{0, 0, 3, 3},
			[][]int{
				{1, 0}, {1, 1}, {0, 1},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			acutalOuput := getNeighbours(test.inputs.x, test.inputs.y, test.inputs.maxX, test.inputs.maxY)
			if !areOutputsEqual(test.expectedOutput, acutalOuput) {
				t.Errorf("expected %v but got %v\n", test.expectedOutput, acutalOuput)
			}
		})
	}
}

func areOutputsEqual(expected, actual [][]int) bool {
	if len(actual) != len(expected) {
		return false
	}
	for index, expectedValue := range expected {
		actualValue := actual[index]
		if actualValue[0] != expectedValue[0] || actualValue[1] != expectedValue[1] {
			return false
		}
	}
	return true
}
