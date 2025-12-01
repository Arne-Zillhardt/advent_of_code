package day0101

import (
	"log"
	"strconv"

	"github.com/Arne-Zillhardt/advent_of_code/2025/fileInput"
)

type Day0101 struct {
	dial *int
}

func NewDay0101() Day0101 {
	dialPosition := 50

	return Day0101{
		dial: &dialPosition,
	}
}
func (d Day0101) Execute(fileInput fileInput.FileInput) {
	listOfOperations := fileInput.GetContentOfFile("./day01/01/input")
	password := 0

	operations := getOperations(listOfOperations)
	amounts := getAmounts(listOfOperations)

	for i := range(len(operations)) {
		zero := false

		if operations[i] == "R" {
			zero = d.increment(amounts[i])
		} else {
			zero = d.decrement(amounts[i])
		}

		if zero {
			password++
		}
	}

	log.Printf("The password is: %d", password)
}

func (d Day0101) increment(amount int) bool {
	*d.dial += amount

	for *d.dial > 99 {
		*d.dial -= 100
	}

	return isZero(*d.dial)
}

func (d Day0101) decrement(amount int) bool {
	*d.dial -= amount

	for *d.dial < 0 {
		*d.dial += 100
	}

	return isZero(*d.dial)
}

func isZero(value int) bool {
	return value == 0
}

func getOperations(listOfOperations []string) []string {
	operations := make([]string, len(listOfOperations))
	for i := range(len(listOfOperations)) {
		operations[i] = listOfOperations[i][:1]
	}

	return operations
}

func getAmounts(listOfAmounts []string) []int {
	amounts := make([]int, len(listOfAmounts))
	for i := range(len(listOfAmounts)) {
		amount := listOfAmounts[i][1:]

		parsedAmount, err := strconv.Atoi(amount)
		if err != nil {
			log.Panic(err)
		}

		amounts[i] = parsedAmount
	}

	return amounts
}
