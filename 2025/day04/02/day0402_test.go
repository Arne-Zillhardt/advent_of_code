package day0402

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

func TestRemoveRoll(t *testing.T) {
	input := []string{
		"@.@@.@@@@@",
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
	expectedOutput := []string{
		"..........",
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

	actualOutput := removeRoll(input, [][]int{{0, 0}, {2, 0}, {3, 0}, {5, 0}, {6, 0}, {7, 0}, {8, 0},{9, 0},})

	if len(expectedOutput) != len(actualOutput) {
		t.Error("Expected the outputs to be equally long")
	}

	for i := range(len(expectedOutput)) {
		if expectedOutput[i] != actualOutput[i] {
			t.Errorf("Expected the outputs to be equal\nExpected: %v\nActual:   %v", expectedOutput, actualOutput)
		}
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

	if numberOfRollsAccessible != 43 {
		t.Errorf("Expected a different number of rolls accesible\nExpected: %d\nActual:   %d", 43, numberOfRollsAccessible)
	}
}
