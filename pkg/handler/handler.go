package handler

import (
	"fmt"

	"github.com/kippophub/go-course/internal/calculator"
)

func ProcessOrder(item string, price float64, qty int) (string, string) {

	msg := fmt.Sprintf("Processing order for %d x %s\n", qty, item)
	total := fmt.Sprintf("Total amount: $%.2f", calculator.CalculateTotal(price, qty))

	return msg, total

}
