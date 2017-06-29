package task3

import (
	"testing"
	"reflect"
)

const EPSILON float64 = 0.00000001

func floatEquals(a, b float64) bool {
	return a-b < EPSILON && b-a < EPSILON
}

var tests = []struct {
	input     Params
	passError bool
	wait      []string
}{
	{
		input: Params{
			[]Triangle{
				{
					Vertices: "ABC",
					A:        10,
					B:        20,
					C:        22.36,
				},
				{
					Vertices: "DEF",
					A:        100,
					B:        200,
					C:        223.6,
				},
				{
					Vertices: "KLM",
					A:        1,
					B:        2,
					C:        2.236,
				},
				{
					Vertices: "OPQ",
					A:        3,
					B:        3,
					C:        3,
				},
			},
		},
		passError: true,
		wait:      []string{"DEF", "ABC", "OPQ", "KLM"},
	},
	{
		input: Params{
			[]Triangle{
				{
					Vertices: "ABC",
					A:        100,
					B:        20,
					C:        22.36,
				},
			},
		},
		passError: false,
	},
}

func TestValidate(t *testing.T) {
	for _, test := range tests {
		if err := Validate(test.input); (err == nil) != test.passError {
			t.Errorf("passError: %v != %v", err == nil, test.passError)
		}
	}
}

func TestRun(t *testing.T) {
	for _, test := range tests {
		if sortedNames, err := Run(test.input); err == nil && !reflect.DeepEqual(sortedNames, test.wait) {
			t.Errorf("sortedNames: %#v != %#v", sortedNames, test.wait)
		}
	}
}

func TestSortTriangles(t *testing.T) {
	for _, test := range tests {
		if err := Validate(test.input); err == nil {
			if sortedNames := SortTriangles(test.input.Triangles); !reflect.DeepEqual(sortedNames, test.wait) {
				t.Errorf("sortedNames: %#v != %#v", sortedNames, test.wait)
			}
		}
	}
}

func TestTriangleValidatePositive(t *testing.T) {
	var triangles = []Triangle{
		{
			Vertices: "ABC",
			A:        22,
			B:        1,
			C:        22.36,
		},
		{
			Vertices: "DEF",
			A:        100,
			B:        200,
			C:        223.6,
		},
		{
			Vertices: "KLM",
			A:        1,
			B:        2,
			C:        2.236,
		},
		{
			Vertices: "OPQ",
			A:        3,
			B:        3,
			C:        3,
		},
	}

	for _, triangle := range triangles {
		if !triangle.Validate() {
			t.Errorf("This triangle must pass validation: %#v", triangle)
		}
	}
}

func TestTriangleValidateNegative(t *testing.T) {
	var triangles = []Triangle{
		{
			Vertices: "ABC",
			A:        -22,
			B:        1,
			C:        22.36,
		},
		{
			Vertices: "DEF",
			A:        1,
			B:        2,
			C:        4,
		},
		{
			Vertices: "KLM",
			A:        0,
			B:        2,
			C:        2.236,
		},
		{
			Vertices: "OPQ",
			A:        3,
			B:        300,
			C:        3,
		},
	}

	for _, triangle := range triangles {
		if triangle.Validate() {
			t.Errorf("This triangle must not pass validation: %#v", triangle)
		}
	}
}

func TestTriangleSquare(t *testing.T) {
	var TrianglesWithAreas = []struct {
		triangle Triangle
		area     float64
	}{
		{triangle: Triangle{Vertices: "ABC", A: 1, B: 1, C: 1}, area: 0.4330127018922193},
		{triangle: Triangle{Vertices: "ABC", A: 1, B: 1, C: 1}, area: 0.433012701892219},
	}

	for _, wait := range TrianglesWithAreas {
		if !floatEquals(wait.triangle.Square(), wait.area) {
			t.Errorf("data.triangle.Square() for triangle %#v is %v must equal %v", wait.triangle, wait.triangle.Square(), wait.area)
		}
	}
}
