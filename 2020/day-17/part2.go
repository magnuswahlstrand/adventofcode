package main

import (
	"bytes"
	"fmt"
)

func solvePart2(size int, input []byte) int {
	mid := size / 2
	inputXY := bytes.Split(input, []byte("\n"))
	offset := len(inputXY) / 2
	wxyzWorld := initWorld2(size, inputXY, mid, offset)

	var active int
	cycles := 6
	for r := 0; r < cycles; r++ {
		prepareNext2(size, wxyzWorld)
		active = applyNext2(size, wxyzWorld)
	}
	return active
}

func applyNext2(size int, wxyzWorld [][][][]*space) int {
	var activeCount int
	for w := 1; w < size-1; w++ {
		for z := 1; z < size-1; z++ {
			for y := 1; y < size-1; y++ {
				for x := 1; x < size-1; x++ {
					if wxyzWorld[w][z][y][x].next == '#' {
						activeCount++
					}

					wxyzWorld[w][z][y][x].current = wxyzWorld[w][z][y][x].next
				}
			}
		}
	}
	return activeCount
}

func prepareNext2(size int, wxyzWorld [][][][]*space) {
	for w := 1; w < size-1; w++ {
		for z := 1; z < size-1; z++ {
			for y := 1; y < size-1; y++ {
				for x := 1; x < size-1; x++ {

					var count int
					for dw := w - 1; dw <= w+1; dw++ {
						for dz := z - 1; dz <= z+1; dz++ {
							for dy := y - 1; dy <= y+1; dy++ {
								for dx := x - 1; dx <= x+1; dx++ {
									// Skip self
									if dw == w && dz == z && dy == y && dx == x {
										continue
									}

									if wxyzWorld[dw][dz][dy][dx].current == '#' {
										count++
									}
								}
							}
						}
					}

					// Check rules
					isActive := wxyzWorld[w][z][y][x].current == '#'
					switch {
					case isActive && (count == 2 || count == 3):
						// Remain active
						wxyzWorld[w][z][y][x].next = '#'
					case !isActive && count == 3:
						// Become active
						wxyzWorld[w][z][y][x].next = '#'
					default:
						// Become/remain inactive
						wxyzWorld[w][z][y][x].next = 0
					}
				}
			}
		}
	}
}

func initWorld2(size int, inputXY [][]byte, mid int, offset int) [][][][]*space {
	wxyzWorld := make([][][][]*space, size)
	for w := 0; w < size; w++ {
		wxyzWorld[w] = make([][][]*space, size)
		for z := 0; z < size; z++ {
			wxyzWorld[w][z] = make([][]*space, size)
			for y := 0; y < size; y++ {
				wxyzWorld[w][z][y] = make([]*space, size)
				for x := 0; x < size; x++ {
					wxyzWorld[w][z][y][x] = &space{current: 0}
				}
			}
		}
	}

	// Add initial active cubes
	for iy := 0; iy < len(inputXY); iy++ {
		for ix := 0; ix < len(inputXY[0]); ix++ {
			if inputXY[iy][ix] == '#' {
				wxyzWorld[mid][mid][iy+mid-offset][ix+mid-offset].current = '#'
			}
		}
	}
	return wxyzWorld
}

func printWorld2(w, z, mid int, offset int, xyzWorld [][][][]*space) {
	for y := mid - offset; y <= mid+offset; y++ {
		for x := mid - offset; x <= mid+offset; x++ {
			switch xyzWorld[w][z][y][x].current {
			case '#':
				fmt.Print("#")
			default:
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
