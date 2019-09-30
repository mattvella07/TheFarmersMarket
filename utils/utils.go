package utils

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

// GetIndexChosen gets input from the user and converts it to an int
func GetIndexChosen(readFrom io.Reader) (int, error) {
	stdin := bufio.NewReader(readFrom)
	choiceStr, err := stdin.ReadString('\n')
	if err != nil {
		return -1, errors.New("Invalid choice, please try again")
	}

	// Remove new line characters
	choiceStr = strings.Replace(choiceStr, "\r", "", -1)
	choiceStr = strings.Replace(choiceStr, "\n", "", -1)

	choice, err := strconv.Atoi(choiceStr)
	if err != nil {
		return -1, errors.New("Invalid choice, please try again")
	}

	// Subtract 1 to use index instead of number displayed
	choice--

	return choice, nil
}

// ChoiceValid makes sure the chosen index is within the range of the passed in
// slice length
func ChoiceValid(choice, sliceLength int) bool {
	if choice >= 0 && choice < sliceLength {
		return true
	}

	return false
}
