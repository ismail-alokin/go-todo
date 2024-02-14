package input_output

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
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

func ReceiveArrayElementId(arr interface{}) (int, error) {
	arrayIdStr := GetAString("Enter number")
	fmt.Println(arrayIdStr)

	arrValue := reflect.ValueOf(arr)
	if arrValue.Kind() != reflect.Array && arrValue.Kind() != reflect.Slice {
		fmt.Println("Input is not an array or a slice")
		return 0, errors.New("invalid type")
	}

	lengthOfArray := arrValue.Len()

	arrayId, err := strconv.Atoi(arrayIdStr)
	if err != nil || arrayId > lengthOfArray {
		fmt.Println("Invalid input")
		return 0, errors.New("invalid input")
	}
	return arrayId - 1, nil
}
