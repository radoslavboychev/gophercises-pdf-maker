package util

import "fmt"

func Currencyfy(input float64) string {
	return fmt.Sprintf("$%.2f", input)
}
