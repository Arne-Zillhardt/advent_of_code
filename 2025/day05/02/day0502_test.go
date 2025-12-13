package day0502

import (
	"testing"
)

var input []string = []string{"3-5", "10-14", "16-20", "12-18", "", "1", "5", "8", "11", "17", "32",}

func TestGetFreshRanges(t *testing.T) {
	expectedOutput := [][]int {
		{3, 5},
		{10, 20},
	}

	actualOutput := getFreshRanges(input)

	compareArrays(t, expectedOutput, actualOutput)
}

func TestNumberOfFreshIDs(t *testing.T) {
	expectedOutput := 14

	actualOutput := numberOfFreshIDs(input)

	if expectedOutput != actualOutput {
		t.Errorf("Expected the output to be equal to the expected output\nExpected: %d\nActual:   %d", expectedOutput, actualOutput)
	}
}

func TestInsertIntoFreshRangesNoMergeSameStart(t *testing.T) {
	input := [][]int {
		{3, 5},
		{10, 14},
		{16, 20},
	}
	expectedOutput := [][]int {
		{10, 14},
		{16, 20},
		{3, 8},
	}
	toInsert := []int {3, 8}
	testInsertIntoFreshRanges(t, input, expectedOutput, toInsert)

	input = [][]int {
		{3, 5},
		{10, 14},
		{16, 20},
	}
	expectedOutput = [][]int {
		{10, 14},
		{16, 20},
		{3, 5},
	}
	toInsert = []int {3, 4}
	testInsertIntoFreshRanges(t, input, expectedOutput, toInsert)
}

func TestInsertIntoFreshRangesNoMergeNoOverlap(t *testing.T) {
	input := [][]int {
		{3, 5},
		{10, 14},
		{16, 20},
	}
	expectedOutput := [][]int {
		{3, 5},
		{10, 14},
		{16, 20},
		{22, 50},
	}
	toInsert := []int {22, 50}
	testInsertIntoFreshRanges(t, input, expectedOutput, toInsert)
}

func TestInsertIntoFreshRangesNoMergeDifferentStart(t *testing.T) {
	input := [][]int {
		{3, 5},
		{10, 14},
		{16, 20},
	}
	expectedOutput := [][]int {
		{10, 14},
		{16, 20},
		{1, 5},
	}
	toInsert := []int {1, 5}
	testInsertIntoFreshRanges(t, input, expectedOutput, toInsert)
}

func TestInsertIntoFreshRangesNoMergeStartNearEnd(t *testing.T) {
	input := [][]int {
		{3, 5},
		{10, 14},
		{16, 20},
	}
	expectedOutput := [][]int {
		{10, 14},
		{16, 20},
		{1, 5},
	}
	toInsert := []int {1, 2}
	testInsertIntoFreshRanges(t, input, expectedOutput, toInsert)
}

func TestInsertIntoFreshRangesNoMergeEndNearStart(t *testing.T) {
	input := [][]int {
		{3, 5},
		{10, 14},
		{16, 20},
	}
	expectedOutput := [][]int {
		{10, 14},
		{16, 20},
		{3, 8},
	}
	toInsert := []int {6, 8}
	testInsertIntoFreshRanges(t, input, expectedOutput, toInsert)
}

func TestInsertIntoFreshRangesMerge(t *testing.T) {
	input := [][]int {
		{10, 14},
		{16, 18},
		{3, 5},
	}
	expectedOutput := [][]int {
		{2, 20},
	}
	toInsert := []int {2, 20}
	testInsertIntoFreshRanges(t, input, expectedOutput, toInsert)
}

func testInsertIntoFreshRanges(t *testing.T, input [][]int, expectedOutput[][]int, toInsert []int) {
	actualOutput := insertIntoFreshRanges(input, toInsert)

	compareArrays(t, actualOutput, expectedOutput)
}

func TestInRange(t *testing.T) {
	first := []int{3, 5}
	second := []int{3, 8}

	isInRange := inRange(first, second)

	if !isInRange {
		t.Error("Expected the ranges to overlap")
	}

	first = []int{3, 4}
	second = []int{6, 8}

	isInRange = inRange(first, second)

	if isInRange {
		t.Error("Expected the ranges to not overlap")
	}

	first = []int{3, 5}
	second = []int{6, 8}

	isInRange = inRange(first, second)

	if !isInRange {
		t.Error("Expected the ranges to overlap")
	}

	first = []int{3, 5}
	second = []int{1, 2}

	isInRange = inRange(first, second)

	if !isInRange {
		t.Error("Expected the ranges to overlap")
	}

	first = []int{3, 5}
	second = []int{2, 20}

	isInRange = inRange(first, second)

	if !isInRange {
		t.Error("Expected the ranges to overlap")
	}
}

func TestRemoveIndex(t *testing.T) {
	//Last index
	input := [][]int {
		{3, 5},
		{10, 14},
		{16, 20},
		{12, 18},
	}
	expectedOutput := [][]int {
		{3, 5},
		{10, 14},
		{16, 20},
	}

	actualOutput := removeIndex(input, 3)

	compareArrays(t, expectedOutput, actualOutput)

	//First index
	input = [][]int {
		{3, 5},
		{10, 14},
		{16, 20},
		{12, 18},
	}
	expectedOutput = [][]int {
		{10, 14},
		{16, 20},
		{12, 18},
	}

	actualOutput = removeIndex(input, 0)

	compareArrays(t, expectedOutput, actualOutput)

	//Middle index
	input = [][]int {
		{3, 5},
		{10, 14},
		{16, 20},
		{12, 18},
	}
	expectedOutput = [][]int {
		{3, 5},
		{16, 20},
		{12, 18},
	}

	actualOutput = removeIndex(input, 1)

	compareArrays(t, expectedOutput, actualOutput)
}

func compareArrays(t *testing.T, expectedOutput [][]int, actualOutput [][]int) {
	if len(expectedOutput) != len(actualOutput) {
		t.Errorf("Expected the output to be equally long as the expected output\nExpected: %v\nActual:   %v", expectedOutput, actualOutput)
		return
	}

	for i := range(len(expectedOutput)){
		if expectedOutput[i][0] != actualOutput[i][0]{
			t.Error("Expected the output to be equal to the expected output")
		}
		if expectedOutput[i][1] != actualOutput[i][1]{
			t.Error("Expected the output to be equal to the expected output")
		}
	}
}
