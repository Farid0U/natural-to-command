package main

import (
	"os"
)

type Config struct {
	APIKey      string
	Model       string
	Interactive bool
	Command     string
}

func NewConfig() *Config {
	config := &Config{
		Model: "openai/gpt-4",
	}
	if apiKey := os.Getenv("OPENROUTER_API_KEY"); apiKey != "" {
		config.APIKey = apiKey
	}
	return config
}
