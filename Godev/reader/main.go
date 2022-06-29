package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a currency: ")
	text2, _ := reader.ReadString('\n')
	text2 = strings.TrimSpace(text2)

	fmt.Print("Enter a number: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	number, _ := strconv.ParseFloat(text, 32)

	fmt.Println("You entered:", text2)
	fmt.Println("You entered:", number)
}
