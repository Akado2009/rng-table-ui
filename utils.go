package main

import (
	"fmt"
	"strconv"
)

func convertAndCheckInt(number string) (int, error) {
	i, err := strconv.Atoi(number)
	if err != nil {
		return 0, err
	}
	if i <= 0 {
		return 0, fmt.Errorf("число не может быть меньше или равно 0")
	}
	return i, nil
}

func convertAndCheckFloat(number string) (float64, error) {
	s, err := strconv.ParseFloat(number, 64)
	return s, err
}

func sliceToString(input [][]float64, precision int) string {
	result := ""
	for i, _ := range input {
		for j, _ := range input[i] {
			template := fmt.Sprintf("\t%%.%df", precision)
			result += fmt.Sprintf(template, input[i][j])
		}
		result += "\n"
	}
	return result
}