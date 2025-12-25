package main

import (
	"fmt"

	"github.com/kippophub/go-course/pkg/handler"
)

func main() {
	fmt.Print(handler.ProcessOrder("Laptop", 1200.50, 1))
	fmt.Println()
	fmt.Print(handler.ProcessOrder("Mouse", 25.00, 2))
}
