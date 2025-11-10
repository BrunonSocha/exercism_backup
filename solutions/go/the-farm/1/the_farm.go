package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
	message string
	numberOfCows int
}

func (i *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", i.numberOfCows, i.message)
}

// TODO: define the 'DivideFood' function
func DivideFood(f FodderCalculator, numberOfCows int) (float64, error) {
	totalAmountOfFodder, err1 := FodderCalculator.FodderAmount(f, numberOfCows)
	if err1 != nil {
		return 0, err1
	}
	multiplier, err2 := FodderCalculator.FatteningFactor(f)
	if err2 != nil {
		return 0, err2
	}
	return totalAmountOfFodder * multiplier/float64(numberOfCows), nil
}


// TODO: define the 'ValidateInputAndDivideFood' function

func ValidateInputAndDivideFood(f FodderCalculator, numberOfCows int) (float64, error) {
	if numberOfCows > 0 {
		return DivideFood(f, numberOfCows)
	}
	return 0, errors.New("invalid number of cows")
}

func ValidateNumberOfCows(numberOfCows int) error {
	if numberOfCows < 0 {
		return &InvalidCowsError{message: "there are no negative cows", numberOfCows: numberOfCows}
	} else if numberOfCows == 0 {
		return &InvalidCowsError{message: "no cows don't need food", numberOfCows: numberOfCows}
	}
	return nil
}


// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
