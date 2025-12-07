package day0302

import (
	"log"
	"math"
	"strconv"

	"github.com/Arne-Zillhardt/advent_of_code/2025/fileInput"
)

type Day0302 struct {
}

func NewDay0302() Day0302 {
	return Day0302{}
}

func (d Day0302) Execute(fileInput fileInput.FileInput) {
	input := fileInput.GetContentOfFile("./day03/01/input")

	sum := 0
	for _, i := range(input) {
		sum += highestJoltage(i)
	}

	log.Printf("The result is: %d", sum)
}

func highestJoltage(batteries string) int {
	splitBatteries := splitBatteries(batteries)
	values := []int{}


	offset := 11
	index := 0
	for range(12) {
		highestBatterie := 0
		indexOfHighestBatterie := 0

		for k, j := range(splitBatteries[index:(len(splitBatteries) - offset)]) {
			if j > highestBatterie {
				highestBatterie = j
				indexOfHighestBatterie = k
			}
		}

		values = append(values, highestBatterie)
		index += indexOfHighestBatterie + 1
		offset--
	}

	result := 0
	for i, v := range(values) {
		result += v * (int(math.Pow10(11 - i)))
	}

	return result
}

func splitBatteries(id string) []int {
	splittedBatteries := []int{}

	for _, i := range(id) {
		value, err := strconv.Atoi(string(i))
		if err != nil {
			log.Panic(err)
		}

		splittedBatteries = append(splittedBatteries, value)
	}

	return splittedBatteries
}
