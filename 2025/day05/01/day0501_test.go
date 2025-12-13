package day0501

import (
	"testing"
)

var input []string = []string{"3-5", "10-14", "16-20", "12-18", "", "1", "5", "8", "11", "17", "32",}

func TestGetFreshRanges(t *testing.T) {
	expectedOutput := [][]int {
		{3, 5},
		{10, 14},
		{16, 20},
		{12, 18},
	}

	actualOutput := getFreshRanges(input)

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

func TestGetIDs(t *testing.T) {
	expectedOutput := []int{1, 5, 8, 11, 17, 32,}

	actualOutput := getIDs(input)

	if len(expectedOutput) != len(actualOutput) {
		t.Errorf("Expected the output to be equally long as the expected output\nExpected: %v\nActual:   %v", expectedOutput, actualOutput)
	}

	for k, v := range(expectedOutput) {
		if v != actualOutput[k] {
			t.Error("Expected the output to be equal to the expected output")
		}
	}
}

func TestNumberOfFreshIDs(t *testing.T) {
	expectedOutput := 3

	actualOutput := numberOfFreshIDs(input)

	if expectedOutput != actualOutput {
		t.Errorf("Expected the output to be equal to the expected output\nExpected: %d\nActual:   %d", expectedOutput, actualOutput)
	}
}
