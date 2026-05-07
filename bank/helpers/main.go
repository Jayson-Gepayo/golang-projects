package helpers

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(value float64, fileName string) {

	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("Failed to read file")
	}

	valueData := string(data)
	value, err := strconv.ParseFloat(valueData, 64)

	if err != nil {
		return 1000, errors.New("failed to parse balance")
	}

	return value, nil
}

func PresentOptions() {
	fmt.Println("What do you want to do")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit")
	fmt.Println("3. Withdraw")
	fmt.Println("4. EXIT")

}
