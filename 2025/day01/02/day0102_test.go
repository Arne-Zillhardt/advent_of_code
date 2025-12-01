package day0102

import (
	"testing"
)

func TestIncrement(t *testing.T) {
	instanceToTest := NewDay0102()

	zero := instanceToTest.increment(1)

	if *instanceToTest.dial != 51 {
		t.Error("Expected to have incremented the dial")
	}

	if zero != 0 {
		t.Error("Expected the dial not to be zero")
	}
}

func TestDecrement(t *testing.T) {
	instanceToTest := NewDay0102()

	zero := instanceToTest.decrement(1)

	if *instanceToTest.dial != 49 {
		t.Error("Expected to have decremented the dial")
	}

	if zero != 0 {
		t.Error("Expected the dial not to be zero")
	}
}

func TestDecrementWithZero(t *testing.T) {
	instanceToTest := NewDay0102()

	zero := instanceToTest.decrement(*instanceToTest.dial)

	if *instanceToTest.dial != 0 {
		t.Error("Expected to have decremented the dial to zero")
	}

	if zero != 1 {
		t.Error("Expected the dial to be zero")
	}
}

func TestIncrementWithOverflow(t *testing.T) {
	instanceToTest := NewDay0102()

	zero := instanceToTest.increment(50)

	if *instanceToTest.dial != 0 {
		t.Error("Expected to have incremented the dial to zero")
	}

	if zero != 1 {
		t.Error("Expected the dial to be zero")
	}
}

func TestDecrementWithUnderflow(t *testing.T) {
	instanceToTest := NewDay0102()

	zero := instanceToTest.decrement(*instanceToTest.dial + 1)

	if *instanceToTest.dial != 99 {
		t.Error("Expected to have decremented the dial to ninetynine")
	}

	if zero != 1 {
		t.Error("Expected the dial to have passed zero one time")
	}
}

func TestIncrementWithMultipleOverflow(t *testing.T) {
	instanceToTest := NewDay0102()

	zero := instanceToTest.increment(1000)

	if *instanceToTest.dial != 50 {
		t.Error("Expected to have incremented the dial to fifty")
	}

	if zero != 10 {
		t.Error("Expected the dial to passed zero ten times")
	}
}

func TestDecrementWithMultipleUnderflow(t *testing.T) {
	instanceToTest := NewDay0102()

	zero := instanceToTest.decrement(1000)

	if *instanceToTest.dial != 50 {
		t.Error("Expected to have decremented the dial to zero")
	}

	if zero != 10 {
		t.Error("Expected the dial to passed zero ten times")
	}
}

func TestGetOperations(t *testing.T) {
	operations := getOperations([]string{"R1", "L1", "R999", "L999"})
	expectedOperations := []string{"R", "L", "R", "L"}

	if len(operations) != 4 {
		t.Error("Expected the operations to be four long")
	}

	for i := range(len(operations)) {
		if operations[i] != expectedOperations[i] {
			t.Error("Expected the operations to be equal")
			break
		}
	}
}

func TestGetAmounts(t *testing.T) {
	amounts := getAmounts([]string{"R1", "L1", "R999", "L999"})
	expectedAmounts := []int{1, 1, 999, 999}

	if len(amounts) != 4 {
		t.Error("Expected the amounts to be four long")
	}

	for i := range(len(amounts)) {
		if amounts[i] != expectedAmounts[i] {
			t.Error("Expected the amounts to be equal")
			break
		}
	}
}

func TestExtended(t *testing.T) {
	instanceToTest := NewDay0102()
	listOfOperations := []string{"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"}
	password := 0

	operations := getOperations(listOfOperations)
	amounts := getAmounts(listOfOperations)

	for i := range(len(operations)) {
		zero := 0

		if operations[i] == "R" {
			zero = instanceToTest.increment(amounts[i])
		} else {
			zero = instanceToTest.decrement(amounts[i])
		}

		password += zero
	}

	if password != 6 {
		t.Error("Expected the password to be 6")
		t.Error(password)
	}

	if *instanceToTest.dial != 32 {
		t.Error("Expected the dial to point at thirtytwo")
	}
}
