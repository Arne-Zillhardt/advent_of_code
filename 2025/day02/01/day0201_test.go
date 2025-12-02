package day0201

import (
	"strings"
	"testing"
)

func FuzzIDValid(f *testing.F) {
	f.Add(1)
	f.Add(12)
	f.Add(100)
	f.Add(115)
	f.Add(222220)
	f.Add(446443)
	f.Add(11204207)
	f.Add(38593856)

	f.Fuzz(func(t *testing.T, id int) {
		isIDValid := idIsValid(id)

		if !isIDValid {
			t.Error("Expected the id to be valid")
		}
	})
}

func TestSplitID(t *testing.T) {
	splitID := splitID(38593856)
	expectedSplitID := []int{3, 8, 5, 9, 3, 8, 5, 6}

	if len(splitID) != len(expectedSplitID) {
		t.Error("Expected the splitIDs to be equally long")
	}

	for i := range(len(splitID)) {
		if splitID[i] != expectedSplitID[i] {
			t.Error("Expected the splitIDs to be equal")
			break
		}
	}
}

func FuzzIDInvalid(f *testing.F) {
	f.Add(22)
	f.Add(99)
	f.Add(222222)
	f.Add(446446)
	f.Add(38593859)

	f.Fuzz(func(t *testing.T, id int) {
		isIDValid := idIsValid(id)

		if isIDValid {
			t.Error("Expected the id to be invalid")
		}
	})
}

func FuzzIDRange(f *testing.F) {
	f.Add("11-22", 11, 22)
	f.Add("118855111880-1188511890", 118855111880, 1188511890)
	f.Add("38593856-38593862", 38593856, 38593862)

	f.Fuzz(func(t *testing.T, input string, expectedStartRange int, expectedEndRange int) {
		startRange, endRange := startAndEndRange(input)

		if startRange != expectedStartRange {
			t.Error("Expected the startRange to be equal to the expectedStartRange")
		}
		if endRange != expectedEndRange {
			t.Error("Expected the endRange to be equal to the expectedEndRange")
		}
	})
}

func TestExtended(t *testing.T) {
	input := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

	addedUp := 0
	for _, idRange := range(strings.Split(input, ",")) {
		startRange, endRange := startAndEndRange(idRange)

		for i := startRange; i <= endRange; i++ {
			if !idIsValid(i) {
				addedUp += i
			}
		}
	}

	if addedUp != 1227775554 {
		t.Errorf("Expected the added up values to be different")
	}
}
