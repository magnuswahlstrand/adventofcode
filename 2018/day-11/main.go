package main

import (
	"fmt"
)

func powerLevel(x, y, sn int) int {
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
	return pl
}

type gridPoint struct {
	x, y, power, width int
}

func (a gridPoint) equals(b gridPoint) bool {
	return a.x == b.x && a.y == b.y && a.power == b.power
}

func createGrid(sn int) [][]int {
	grid := make([][]int, 300)
	for i := range grid {
		grid[i] = make([]int, 300)
	}

	for y := range grid {
		for x := range grid[y] {
			grid[y][x] = powerLevel(x+1, y+1, sn)
		}
	}
	return grid
}

func highestPowerLevelV2(sn int) gridPoint {
	grid := createGrid(sn)

	max := gridPoint{
		power: -100,
	}
	for w := 1; w <= 300; w++ {
		fmt.Println("Searching at:", w, ". Current max is:", max)
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

func highestInGrid(grid [][]int, width int) gridPoint {
	lenX := len(grid[0])
	var max = -1
	var highest gridPoint
	for y := range grid {
	Inner:
		for x := range grid[y] {

			// Investigate this point
			var total int
			for dy := 0; dy < width; dy++ {
				for dx := 0; dx < width; dx++ {

					if y+dy >= len(grid) || x+dx >= lenX {
						continue Inner
					}

					total += grid[y+dy][x+dx]
				}
			}
			if x >= 20 && x < 50 && y >= 40 && y < 79 {

				// if (x == 21 && y == 61) || (x == 33 && y == 45) {
				// 	fmt.Printf("{%d}", grid[y][x])
				// } else {
				// 	fmt.Printf("% 3d", grid[y][x])
				// }
			}
			// fmt.Printf("% 3d (% 3d)", grid[y][x], total)

			if total > max {
				max = total
				highest = gridPoint{
					x:     x,
					y:     y,
					power: total,
					width: width,
				}

			}
			// fmt.Printf("% 3d", total)

		}

		// if y >= 40 && y < 79 {
		// 	fmt.Println("")
		// }
	}
	return highest
}

func main() {
	p := highestPowerLevel(8141)
	fmt.Printf("Highest Power level is: %d at (x=%d,y=%d)", p.power, p.x+1, p.y+1)

	p2 := highestPowerLevelV2(8141)
	fmt.Printf("Highest Power level is: %d at (%d,%d,%d)", p2.power, p2.x+1, p2.y+1, p2.width)

	// printGrid(createGrid(18), 33-1, 45-1)
	// printGrid(createGrid(42), 21-1, 61-1)
	// printGrid(createGrid(8141), point.x, point.y)
}

func printGrid(grid [][]int, x, y int) {
	fmt.Println("")
	for dy := 0; dy < 10; dy++ {
		for dx := 0; dx < 10; dx++ {
			fmt.Printf("% 3d", grid[y+dy-1][x+dx-1])
		}
		fmt.Println("")
	}
}
