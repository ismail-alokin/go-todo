package input_output

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func PrintTitle(title string) {
	fmt.Println(title)
	fmt.Print("-------------------\n\n")
}

func GetAString(message string) string {
	fmt.Println(message)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Println("ERROR:", err)
	}

	return scanner.Text()
}

func ShouldExit(input string) bool {
	return input == "5"
}

func ValidateInput(input string) bool {
	input = strings.TrimSpace(input)
	num, err := strconv.Atoi(input)

	if isInValidInput(input, num, err) {
		log.Println("Invalid input")
		return false
	}
	return true
}

func isInValidInput(input string, num int, err error) bool {
	return (err != nil) || (input == "") || (len(input) > 1) || (num < 1) || (num > 5)
}

// func GetStringArrayOfInputValues(message string, separator string) []string {
// 	fmt.Print(message)

// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	err := scanner.Err()
// 	if err != nil {
// 		log.Println("ERROR:", err)
// 	}

// 	inputString := scanner.Text()

// 	inputArray := strings.Split(inputString, separator)

// 	filteredInputArray := make([]string, 0, len(inputArray))

// 	for _, str := range inputArray {
// 		str = strings.TrimSpace(str)
// 		if str == "" {
// 			continue
// 		}
// 		filteredInputArray = append(filteredInputArray, str)
// 	}

// 	return filteredInputArray
// }

// func GetMultiLineInputStrings(message string) []string {
// 	fmt.Println(message)

// 	var lines []string
// 	scanner := bufio.NewScanner(os.Stdin)

// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		line = strings.TrimSpace(line)

// 		if line == "" {
// 			break
// 		}

// 		lines = append(lines, line)
// 	}

// 	return lines
// }
