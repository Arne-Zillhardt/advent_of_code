package day0201

import (
	"log"
	"strconv"
	"strings"

	"github.com/Arne-Zillhardt/advent_of_code/2025/fileInput"
)

type Day0201 struct {
}

func NewDay0201() Day0201 {
	return Day0201{}
}

func (d Day0201) Execute(fileInput fileInput.FileInput) {
	input := fileInput.GetContentOfFile("./day02/01/input")[0]

	addedUp := 0
	for _, idRange := range(strings.Split(input, ",")) {
		startRange, endRange := startAndEndRange(idRange)

		for i := startRange; i <= endRange; i++ {
			if !idIsValid(i) {
				addedUp += i
			}
		}
	}

	log.Printf("The result is: %d", addedUp)
}

func idIsValid(id int) bool {
	splittedID := splitID(id)

	if len(splittedID) % 2 != 0 {
		return true
	}

	firstHalf := splittedID[:len(splittedID) / 2]
	secondHalf := splittedID[len(splittedID) / 2:]

	equal := true
	for i := range(firstHalf) {
		if firstHalf[i] != secondHalf[i] {
			equal = false
			break
		}
	}

	return !equal
}

func splitID(id int) []int {
	splittedID := []int{}

	factor := 1
	for factor < id {
		factor *= 10
	}

	if factor >= 10 {
		factor /= 10
	}

	for factor != 1 {
		splittedID = append(splittedID, id / factor)

		id = id % factor
		factor /= 10

	}

	splittedID = append(splittedID, id)

	return splittedID
}

func startAndEndRange(ranges string) (int, int) {
	startRange := strings.Split(ranges, "-")[0]
	endRange := strings.Split(ranges, "-")[1]

	parsedStartRange, err := strconv.Atoi(startRange)
	if err != nil {
		log.Panic(err)
	}

	parsedEndRange, err := strconv.Atoi(endRange)
	if err != nil {
		log.Panic(err)
	}

	return parsedStartRange, parsedEndRange
}
