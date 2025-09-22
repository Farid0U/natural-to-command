package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInteractiveInput() string {
	fmt.Print("Enter command: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	config := NewConfig()
	parseFlags(config)

	client := NewOpenRouterClient(config.APIKey, config.Model)
	executor := NewCommandExecutor()

	humanCommand := config.Command
	if config.Interactive {
		humanCommand = getInteractiveInput()
	}

	linuxCommand, err := client.TranslateCommand(humanCommand)
	if err != nil {
		fmt.Println("Error translating command:", err)
		os.Exit(1)
	}

	if !executor.IsSafe(linuxCommand) {
		fmt.Println("Command is not safe")
		os.Exit(1)
	}

	risk := executor.AssessRisk(linuxCommand)
	riskStr := "LOW"
	if risk == Medium {
		riskStr = "MEDIUM"
	} else if risk == High {
		riskStr = "HIGH"
	}

	fmt.Printf("Generated command: %s\n", linuxCommand)
	fmt.Printf("Risk level: %s\n", riskStr)
	fmt.Print("Execute? (y/N): ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	confirm := strings.ToLower(strings.TrimSpace(scanner.Text()))

	if confirm == "y" || confirm == "yes" {
		err := executor.Execute(linuxCommand)
		if err != nil {
			fmt.Println("Error executing command:", err)
		}
	}
}
