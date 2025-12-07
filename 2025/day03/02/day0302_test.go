package day0302

import (
	"testing"
)

func FuzzHighestJoltage(f *testing.F) {
	f.Add("987654321111111", 987654321111)
	f.Add("811111111111119", 811111111119)
	f.Add("234234234234278", 434234234278)
	f.Add("818181911112111", 888911112111)

	f.Fuzz(func(t *testing.T, input string, expectedOutput int) {
		output := highestJoltage(input)

		if output != expectedOutput {
			t.Errorf("Expected the output to be equal to the expected output\nExpected: %d\nActual:   %d", expectedOutput, output)
		}
	})
}

func TestSplitID(t *testing.T) {
	splitBatteries := splitBatteries("38593856")
	expectedSplitID := []int{3, 8, 5, 9, 3, 8, 5, 6}

	if len(splitBatteries) != len(expectedSplitID) {
		t.Error("Expected the splitBatteries to be equally long")
	}

	for i := range(len(splitBatteries)) {
		if splitBatteries[i] != expectedSplitID[i] {
			t.Error("Expected the splitBatteries to be equal")
			break
		}
	}
}

func TestExtended(t *testing.T) {
	input := []string{"987654321111111", "811111111111119", "234234234234278", "818181911112111"}

	sum := 0
	for _, i := range(input) {
		sum += highestJoltage(i)
	}

	if sum != 3121910778619 {
		t.Error("Expected a different value for sum")
	}
}
