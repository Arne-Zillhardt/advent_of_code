package day0401

import (
	"testing"
)

func TestFieldContainsPaper(t *testing.T) {
	input := []string{
		"..@@.@@@@.",
		"@@@.@.@.@@",
		"@@@@@.@.@@",
		"@.@@@@..@.",
		"@@.@@@@.@@",
		".@@@@@@@.@",
		".@.@.@.@@@",
		"@.@@@.@@@@",
		".@@@@@@@@.",
		"@.@.@@@.@.",
	}
	

	testFieldContainsPaperNegative(t, input, 0, 0);
	testFieldContainsPaperNegative(t, input, -1, -1);
	testFieldContainsPaperPositive(t, input, 2, 0);
	testFieldContainsPaperPositive(t, input, 2, 1);
}

func testFieldContainsPaperNegative(t *testing.T, input []string, x int, y int) {
	containsPaper := fieldContainsPaper(input, x, y)

	if containsPaper {
		t.Error("Didn't expect the field to contain paper")
	}
}

func testFieldContainsPaperPositive(t *testing.T, input []string, x int, y int) {
	containsPaper := fieldContainsPaper(input, x, y)

	if !containsPaper {
		t.Error("Did expect the field to contain paper")
	}
}

func TestRollsAccessible(t *testing.T) {
	input := []string{
		"..@@.@@@@.",
		"@@@.@.@.@@",
		"@@@@@.@.@@",
		"@.@@@@..@.",
		"@@.@@@@.@@",
		".@@@@@@@.@",
		".@.@.@.@@@",
		"@.@@@.@@@@",
		".@@@@@@@@.",
		"@.@.@@@.@.",
	}

	numberOfRollsAccessible := rollsAccessible(input)

	if numberOfRollsAccessible != 13 {
		t.Errorf("Expected a different number of rolls accesible\nExpected: %d\nActual:   %d", 13, numberOfRollsAccessible)
	}
}
