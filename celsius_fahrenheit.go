package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Укажите температуру в Цельсиях")
		return
	}

	celsiusStr := os.Args[1]
	celsius, err := strconv.ParseFloat(celsiusStr, 64)
	if err != nil {
		fmt.Println("Ошибка: введено не число")
		return
	}

	fahrenheit := celsius*9/5 + 32
	fmt.Printf("%.1f °F\n", fahrenheit)
}

