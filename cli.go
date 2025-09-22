package main

import (
	"flag"
	"fmt"
	"os"
)

func parseFlags(config *Config) {
	flag.StringVar(&config.Command, "command", "", "Human language command to translate")
	flag.BoolVar(&config.Interactive, "interactive", false, "Interactive mode")
	flag.StringVar(&config.Model, "model", config.Model, "AI model to use")
	flag.Parse()

	if config.APIKey == "" {
		fmt.Println("OPENROUTER_API_KEY environment variable not set")
		os.Exit(1)
	}
}
