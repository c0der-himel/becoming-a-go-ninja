package main

import (
	"bufio"
	"cli-calculator/internal/calc"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func main() {
	green := color.New(color.FgGreen).SprintFunc()
	boldGreen := color.New(color.FgHiGreen, color.Bold).SprintFunc()

	fmt.Println(green("======================================="))
	fmt.Println(boldGreen("        🧮 RETRO GO CALCULATOR 🧮"))
	fmt.Println(green("======================================="))
	fmt.Println(green("Type 'exit' to quit"))
	fmt.Println(green("Format: 10+5"))
	fmt.Println(green("---------------------------------------"))

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(boldGreen("calc> "))
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println(green("Goodbye, Himel 👾"))
			break
		}

		result, err := calc.Evaluate(input)
		if err != nil {
			fmt.Println(color.RedString("ERROR:"), err)
			continue
		}

		fmt.Println(green("= "), boldGreen(fmt.Sprintf("%.2f", result)))
	}
}
