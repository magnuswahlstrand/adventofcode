package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRows(t *testing.T) {
	row := parseRow("BFFFBBFRRR")
	require.Equal(t, 70, row)
	row = parseRow("FBFBBFFRLR")
	require.Equal(t, 44, row)
	row = parseRow("FFFBBBFRRR")
	require.Equal(t, 14, row)
	row = parseRow("BBFFBBFRLL")
	require.Equal(t, 102, row)
}
func TestCols(t *testing.T) {
	col := parseCols("FBFBBFFRLR")
	require.Equal(t, 5, col)
	col = parseCols("BFFFBBFRRR")
	require.Equal(t, 7, col)
	col = parseCols("FFFBBBFRRR")
	require.Equal(t, 7, col)
	col = parseCols("BBFFBBFRLL")
	require.Equal(t, 4, col)
}

func TestSeatID(t *testing.T) {
	seatID := parseSeatID("BFFFBBFRRR")
	require.Equal(t, 567, seatID)
	seatID = parseSeatID("FFFBBBFRRR")
	require.Equal(t, 119, seatID)
	seatID = parseSeatID("BBFFBBFRLL")
	require.Equal(t, 820, seatID)
}