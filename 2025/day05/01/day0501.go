package day0501

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/Arne-Zillhardt/advent_of_code/2025/fileInput"
)

type Day0501 struct {
}

func NewDay0501() Day0501 {
	return Day0501{}
}

func (d Day0501)Execute(fileInput fileInput.FileInput) {
	input := fileInput.GetContentOfFile("./day05/01/input")
	freshIDs := numberOfFreshIDs(input)

	log.Printf("The result is: %d", freshIDs)
}

func numberOfFreshIDs(input []string) int {
	numberOfFreshIDs := 0

	for _, id := range(getIDs(input)) {
		for _, freshRange := range(getFreshRanges(input)) {
			if rangeContainsID(freshRange, id) {
				numberOfFreshIDs++
				break
			}
		}
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

		freshRanges = append(freshRanges, []int{startRange, endRange})
	}

	return freshRanges
}

func getIDs(input []string) []int {
	ids := make([]int, 0)

	for _, v := range(input) {
		isID, err := regexp.MatchString(`^\d+$`, v)
		if err != nil {
			log.Panic(err)
		}

		if !isID {
			continue
		}

		id, err := strconv.Atoi(v)
		if err != nil {
			log.Panic(err)
		}

		ids = append(ids, id)
	}

	return ids
}

func rangeContainsID(freshRange []int, id int) bool {
	if id < freshRange[0] {
		return false
	}
	if id > freshRange[1] {
		return false
	}

	return true
}
