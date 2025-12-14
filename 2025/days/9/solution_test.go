package main

import (
	"testing"
)

func TestFourCorners(t *testing.T) {
	tests := []struct {
		testName       string
		point1         Point
		point2         Point
		expectedOutput []Point
	}{
		{
			"y = -x, 1",
			Point{2, 2},
			Point{5, 5},
			[]Point{{2, 2}, {5, 2}, {5, 5}, {2, 5}},
		},
		{
			"y = -x, 2",
			Point{5, 5},
			Point{2, 2},
			[]Point{{2, 2}, {5, 2}, {5, 5}, {2, 5}},
		},
		{
			"y = x, 1",
			Point{2, 5},
			Point{5, 2},
			[]Point{{2, 2}, {5, 2}, {5, 5}, {2, 5}},
		},
		{
			"y = x, 2",
			Point{5, 2},
			Point{2, 5},
			[]Point{{2, 2}, {5, 2}, {5, 5}, {2, 5}},
		},
		{
			"horizontal, 1",
			Point{2, 2},
			Point{5, 2},
			[]Point{{2, 2}, {5, 2}, {5, 2}, {2, 2}},
		},
		{
			"horizontal, 2",
			Point{5, 2},
			Point{2, 2},
			[]Point{{2, 2}, {5, 2}, {5, 2}, {2, 2}},
		},
		{
			"vertical, 1",
			Point{2, 2},
			Point{2, 5},
			[]Point{{2, 2}, {2, 2}, {2, 5}, {2, 5}},
		},
		{
			"vertical, 2",
			Point{2, 5},
			Point{2, 2},
			[]Point{{2, 2}, {2, 2}, {2, 5}, {2, 5}},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			acutalOuput := getFourCorners(test.point1, test.point2)
			for index, elem1 := range test.expectedOutput {
				if elem1 != acutalOuput[index] {
					t.Errorf("expected %v but got %v\n", test.expectedOutput, acutalOuput)
				}
			}
		})
	}
}
