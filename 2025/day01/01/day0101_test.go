package day0101

import (
	"testing"
)

func TestIncrement(t *testing.T) {
	instanceToTest := NewDay0101()

	zero := instanceToTest.increment(1)

	if *instanceToTest.dial != 51 {
		t.Error("Expected to have incremented the dial")
	}

	if zero {
		t.Error("Expected the dial not to be zero")
	}
}

func TestDecrement(t *testing.T) {
	instanceToTest := NewDay0101()

	zero := instanceToTest.decrement(1)

	if *instanceToTest.dial != 49 {
		t.Error("Expected to have decremented the dial")
	}

	if zero {
		t.Error("Expected the dial not to be zero")
	}
}

func TestDecrementWithZero(t *testing.T) {
	instanceToTest := NewDay0101()

	zero := instanceToTest.decrement(*instanceToTest.dial)

	if *instanceToTest.dial != 0 {
		t.Error("Expected to have decremented the dial to zero")
	}

	if !zero {
		t.Error("Expected the dial to be zero")
	}
}

func TestIncrementWithOverflow(t *testing.T) {
	instanceToTest := NewDay0101()

	zero := instanceToTest.increment(50)

	if *instanceToTest.dial != 0 {
		t.Error("Expected to have incremented the dial to zero")
	}

	if !zero {
		t.Error("Expected the dial to be zero")
	}
}

func TestDecrementWithUnderflow(t *testing.T) {
	instanceToTest := NewDay0101()

	zero := instanceToTest.decrement(*instanceToTest.dial + 1)

	if *instanceToTest.dial != 99 {
		t.Error("Expected to have decremented the dial to ninetynine")
	}

	if zero {
		t.Error("Expected the dial not to be zero")
	}
}

func TestIncrementWithMultipleOverflow(t *testing.T) {
	instanceToTest := NewDay0101()

	zero := instanceToTest.increment(150)

	if *instanceToTest.dial != 0 {
		t.Error("Expected to have incremented the dial to zero")
	}

	if !zero {
		t.Error("Expected the dial to be zero")
	}
}

func TestDecrementWithMultipleUnderflow(t *testing.T) {
	instanceToTest := NewDay0101()

	zero := instanceToTest.decrement(151)

	if *instanceToTest.dial != 99 {
		t.Error("Expected to have decremented the dial to ninetynine")
	}

	if zero {
		t.Error("Expected the dial not to be zero")
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
