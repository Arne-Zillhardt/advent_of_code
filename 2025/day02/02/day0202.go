package day0202

import (
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/Arne-Zillhardt/advent_of_code/2025/fileInput"
)

type Day0202 struct {
}

func NewDay0202() Day0202 {
	return Day0202{}
}

func (d Day0202) Execute(fileInput fileInput.FileInput) {
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

	pattern := []int{}
	for _, i := range(splittedID) {
		lenOfSplittedIDs := float64(len(splittedID))
		half := math.Floor(lenOfSplittedIDs / 2.0)
		pattern = append(pattern, i)

		if len(pattern) > int(half) {
			return true
		}

		if listsEqual(pattern, splittedID) {
			return false
		}
	}

	return true
}

func splitID(id int) []int {
	splittedID := []int{}

	factor := 1
	for factor < id + 1 {
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

func listsEqual(list01 []int, list02 []int) bool {
	if len(list02) % len(list01) != 0 {
		return false
	}

	for i := 0; i < len(list02); i += len(list01) {
		currentList := list02[i:i + len(list01)]

		for j := range(currentList) {
			if list01[j] != currentList[j] {
				return false
			}
		}
	}

	return true
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
