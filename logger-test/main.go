package main

import (
	"fmt"
	"github.com/qazaq159/awesome/logger"
	"github.com/qazaq159/awesome/math"
)

func main() {
	logger.Log("hey there")
	fmt.Println(math.Sum(7, 5))
	fmt.Println(math.Multiple(5, 11))
}
