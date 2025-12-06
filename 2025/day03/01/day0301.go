package day0301

import (
	"log"
	"strconv"

	"github.com/Arne-Zillhardt/advent_of_code/2025/fileInput"
)

type Day0301 struct {
}

func NewDay0301() Day0301 {
	return Day0301{}
}

func (d Day0301) Execute(fileInput fileInput.FileInput) {
	input := fileInput.GetContentOfFile("./day03/01/input")

	sum := 0
	for _, i := range(input) {
		sum += highestJoltage(i)
	}

	log.Printf("The result is: %d", sum)
}

func highestJoltage(batteries string) int {
	splitBatteries := splitBatteries(batteries)

	highestBatterie := 0
	indexOfHighestBatterie := 0

	for i := 0; i <= len(splitBatteries) - 2; i++ {
		if highestBatterie < splitBatteries[i] {
			highestBatterie = splitBatteries[i]
			indexOfHighestBatterie = i 
		}
	}

	secondHighestBatterie := 0
	for _, i := range(splitBatteries[indexOfHighestBatterie + 1:]) {
		if secondHighestBatterie < i {
			secondHighestBatterie = i
		}
	}

	return (highestBatterie * 10) + secondHighestBatterie
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
