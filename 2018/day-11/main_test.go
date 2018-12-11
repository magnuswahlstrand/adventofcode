package main

import (
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestPowerLevel(t *testing.T) {
	tcs := []struct {
		name     string
		x, y, sn int
		expected float64
	}{

		{name: "small", x: 3, y: 5, sn: 8, expected: 4},
		{name: "negative", x: 122, y: 79, sn: 57, expected: -5},
		{name: "zero", x: 217, y: 196, sn: 39, expected: 0},
		{name: "postive", x: 101, y: 153, sn: 71, expected: 4},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {

			res := powerLevel(tc.x, tc.y, tc.sn)

			if res != tc.expected {
				t.Fatalf("\nExpected %f, got %f\n", tc.expected, res)
			}
		})
	}
}

func TestHighestPowerLevel(t *testing.T) {
	tcs := []struct {
		name     string
		sn       int
		expected gridPoint
	}{

		{name: "sn18", sn: 18, expected: gridPoint{x: 33 - 1, y: 45 - 1, power: 29, width: 3}},
		{name: "sn42", sn: 42, expected: gridPoint{x: 21 - 1, y: 61 - 1, power: 30, width: 3}},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {

			res := highestPowerLevel(tc.sn)

			if res != tc.expected {
				t.Fatalf("\nExpected %s, got %s\n", tc.expected, res)
			}
		})
	}
}

// func TestHighestPowerLevelV2(t *testing.T) {
// 	tcs := []struct {
// 		name     string
// 		sn       int
// 		expected gridPoint
// 	}{

// 		{name: "sn18", sn: 18, expected: gridPoint{x: 90 - 1, y: 269 - 1, power: 113, width: 16}},
// 		{name: "sn42", sn: 42, expected: gridPoint{x: 232 - 1, y: 251 - 1, power: 119, width: 12}},
// 	}

// 	for _, tc := range tcs {
// 		t.Run(tc.name, func(t *testing.T) {

// 			res := highestPowerLevelV2(tc.sn)

// 			if res != tc.expected {
// 				t.Fatalf("\nExpected %d, got %d\n", tc.expected, res)
// 			}
// 		})
// 	}
// }

func TestHighestInGrid(t *testing.T) {
	tcs := []struct {
		name     string
		grid     *mat.Dense
		expected gridPoint
	}{

		{name: "sn18", grid: mat.NewDense(5, 5, []float64{
			-2, -4, 4, 4, 4,
			-4, 4, 4, 4, -5,
			4, 3, 3, 4, -4,
			1, 1, 2, 4, -3,
			-1, 0, 2, -5, -2,
		}), expected: gridPoint{
			x:     1,
			y:     1,
			power: 29,
		}},
		{name: "sn42", grid: mat.NewDense(5, 5, []float64{
			-3, 4, 2, 2, 2,
			-4, 4, 3, 3, 4,
			-5, 3, 3, 4, -4,
			4, 3, 3, 4, -3,
			3, 3, 3, -5, -1,
		}), expected: gridPoint{
			x:     1,
			y:     1,
			power: 30,
		}},
		{name: "snMagnus", grid: mat.NewDense(5, 5,
			[]float64{-3, 4, 2, 2, 2,
				-4, 4, 3, 3, 4,
				-5, 3, 3, 4, -4,
				4, 3, 3, 4, -15,
				3, 3, 3, -5, 30}), expected: gridPoint{
			x:     1,
			y:     1,
			power: 30,
		}},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {

			res := highestInGrid(tc.grid, 3)

			if !res.equals(tc.expected) {
				t.Fatalf("\nExpected %s, got %s\n", tc.expected, res)
			}
		})
	}
}
