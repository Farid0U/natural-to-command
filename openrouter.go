package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"
)

type OpenRouterClient struct {
	apiKey string
	model  string
}

type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatResponse struct {
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Message Message `json:"message"`
}

func NewOpenRouterClient(apiKey, model string) *OpenRouterClient {
	return &OpenRouterClient{apiKey: apiKey, model: model}
}

func (c *OpenRouterClient) TranslateCommand(humanCommand string) (string, error) {
	os := "Linux"
	if runtime.GOOS == "windows" {
		os = "Windows"
	}
	prompt := fmt.Sprintf("Translate this human command to a safe %s command: '%s'. Output only the command, no explanation.", os, humanCommand)
	return c.makeRequest(prompt)
}

func (c *OpenRouterClient) makeRequest(prompt string) (string, error) {
	url := "https://openrouter.ai/api/v1/chat/completions"
	reqBody := ChatRequest{
		Model: c.model,
		Messages: []Message{
			{Role: "user", Content: prompt},
		},
	}
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var chatResp ChatResponse
	if err := json.Unmarshal(body, &chatResp); err != nil {
		return "", err
	}

	if len(chatResp.Choices) == 0 {
		return "", fmt.Errorf("no choices in response")
	}

	return chatResp.Choices[0].Message.Content, nil
}
