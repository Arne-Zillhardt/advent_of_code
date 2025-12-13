package day0502

import (
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/Arne-Zillhardt/advent_of_code/2025/fileInput"
)

type Day0502 struct {
}

func NewDay0502() Day0502 {
	return Day0502{}
}

func (d Day0502)Execute(fileInput fileInput.FileInput) {
	input := fileInput.GetContentOfFile("./day05/01/input")
	freshIDs := numberOfFreshIDs(input)

	log.Printf("The result is: %d", freshIDs)
}

func numberOfFreshIDs(input []string) int {
	numberOfFreshIDs := 0
	for _, freshRange := range(getFreshRanges(input)) {
		numberOfFreshIDs += (freshRange[1] - freshRange[0]) + 1
	}
	return numberOfFreshIDs
}

func getFreshRanges(input []string) [][]int {
	freshRanges := make([][]int, 0)

	for _, v := range(input) {
		if v == "" {
			break
		}

		start, end, found := strings.Cut(v, "-")
		if !found {
			log.Panic("Expected a range: ", v)
		}

		startRange, err := strconv.Atoi(start)
		if err != nil {
			log.Panic(err)
		}

		endRange, err := strconv.Atoi(end)
		if err != nil {
			log.Panic(err)
		}

		freshRanges = insertIntoFreshRanges(freshRanges, []int{startRange, endRange})
	}

	return freshRanges
}

func insertIntoFreshRanges(input [][]int, toInsert []int) [][]int {
	changedInsert := false
	for i, v := range(input) {
		if inRange(toInsert, v){
			input = removeIndex(input, i)

			newStart := int(math.Min(float64(toInsert[0]), float64(v[0])))
			newEnd := int(math.Max(float64(toInsert[1]), float64(v[1])))
			toInsert = []int{newStart, newEnd}
			changedInsert = true
		}
	}

	if !changedInsert {
		input = append(input, toInsert)
		return input
	}

	return insertIntoFreshRanges(input, toInsert)
}

func inRange(first []int, second []int) bool {
	if first[0] > (second[1] + 1) {
		return false
	}

	if first[1] < (second[0] - 1) {
		return false
	}


	return true
}

func removeIndex(array [][]int, index int) [][]int {
	firstHalf := make([][]int, 0)
	firstHalf = array[0:index]

	secondHalf := make([][]int, 0)
	if index < len(array) - 1 {
		secondHalf = array[index + 1:len(array)]
	}

	return append(firstHalf, secondHalf...)
}
