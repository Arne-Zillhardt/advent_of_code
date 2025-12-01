package day0102

import (
	"log"
	"strconv"

	"github.com/Arne-Zillhardt/advent_of_code/2025/fileInput"
)

type Day0102 struct {
	dial *int
}

func NewDay0102() Day0102 {
	dialPosition := 50

	return Day0102{
		dial: &dialPosition,
	}
}
func (d Day0102) Execute(fileInput fileInput.FileInput) {
	listOfOperations := fileInput.GetContentOfFile("./day01/01/input")
	password := 0

	operations := getOperations(listOfOperations)
	amounts := getAmounts(listOfOperations)

	for i := range(len(operations)) {
		zero := 0

		if operations[i] == "R" {
			zero = d.increment(amounts[i])
		} else {
			zero = d.decrement(amounts[i])
		}

		password += zero
	}

	log.Printf("The password is: %d", password)
}

func (d Day0102) increment(amount int) int {
	zero := 0

	for range(amount) {
		*d.dial++

		if *d.dial > 99 {
			*d.dial = 0
		}

		if *d.dial == 0 {
			zero++
		}
	}

	return zero
}

func (d Day0102) decrement(amount int) int {
	zero := 0

	for range(amount) {
		*d.dial--

		if *d.dial < 0 {
			*d.dial = 99
		}

		if *d.dial == 0 {
			zero++
		}
	}

	return zero
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
