package main

import (
	"fmt"

	"github.com/kippophub/go-course/pkg/utils"
)

func main() {
	b := utils.CountChars("Привет!")
	fmt.Printf("%d", b)
}
