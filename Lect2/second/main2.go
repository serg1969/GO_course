package main

import (
	"fmt"
	"strings"
)

//K solution

func main() {
	var total string = "чат"
	var message string

	fmt.Scan(&message)

	if strings.Contains(message, total) {
		fmt.Println("БОТ")
	} else {
		fmt.Println("НЕ БОТ")
	}
}
