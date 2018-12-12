package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func powerLevel(x, y, sn int) float64 {
	// Find the fuel cell's rack ID, which is its X coordinate plus 10.
	rackID := x + 10

	// Begin with a power level of the rack ID times the Y coordinate.
	pl := rackID * y

	// Increase the power level by the value of the grid serial number (your puzzle input).
	pl += sn

	// Set the power level to itself multiplied by the rack ID.
	pl *= rackID

	// Keep only the hundreds digit of the power level (so 12345 becomes 3; numbers with no hundreds digit become 0).
	pl = (pl / 100) % 10

	// Subtract 5 from the power level.
	pl -= 5
	return float64(pl)
}

type gridPoint struct {
	x, y  int
	power float64
	width int
}

func (a gridPoint) String() string {
	return fmt.Sprintf("% 4d,% 4d,% 4.0f", a.x+1, a.y+1, a.power)
}

func (a gridPoint) equals(b gridPoint) bool {
	return a.x == b.x && a.y == b.y && a.power == b.power
}

func createGrid(sn int) *mat.Dense {
	g := mat.NewDense(300, 300, nil)

	for y := 0; y < 300; y++ {
		for x := 0; x < 300; x++ {
			g.Set(y, x, powerLevel(x+1, y+1, sn))
		}
	}

	return g
}

func highestPowerLevelV2(sn int) gridPoint {
	grid := createGrid(sn)

	max := gridPoint{
		power: -100,
	}
	for w := 1; w <= 300; w++ {
		fmt.Printf("Search width % 3d, and current max is%s\n", w, max)
		p := highestInGrid(grid, w)

		if p.power > max.power {
			max = p
		}
	}
	return max
}

func highestPowerLevel(sn int) gridPoint {
	grid := createGrid(sn)
	return highestInGrid(grid, 3)
}

func highestInGrid(grid *mat.Dense, width int) gridPoint {
	lenY, lenX := grid.Dims()

	var max float64 = -1
	var highest gridPoint
	for y := 0; y < lenY && y+width <= lenY; y++ {
		for x := 0; x < lenX && x+width <= lenX; x++ {

			sl := grid.Slice(y, y+width, x, x+width)
			total := mat.Sum(sl)

			if total > max {
				max = total
				highest = gridPoint{
					x:     x,
					y:     y,
					power: total,
					width: width,
				}
			}
		}
	}
	return highest
}

func main() {
	p := highestPowerLevel(8141)
	fmt.Printf("Highest Power level is: %f at (x=%d,y=%d)\n", p.power, p.x+1, p.y+1)

	p2 := highestPowerLevelV2(8141)
	fmt.Printf("Highest Power level is: %f at (%s)\n", p2.power, p2)
}
