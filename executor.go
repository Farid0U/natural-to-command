package main

import (
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

type RiskLevel int

const (
	Low RiskLevel = iota
	Medium
	High
)

type CommandExecutor struct{}

func NewCommandExecutor() *CommandExecutor {
	return &CommandExecutor{}
}

func (e *CommandExecutor) Execute(command string) error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", command)
	} else {
		cmd = exec.Command("sh", "-c", command)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (e *CommandExecutor) AssessRisk(command string) RiskLevel {
	if strings.Contains(command, "rm") || strings.Contains(command, "dd") {
		return High
	}
	if strings.Contains(command, "sudo") {
		return Medium
	}
	return Low
}

func (e *CommandExecutor) IsSafe(command string) bool {
	dangerous := regexp.MustCompile(`rm\s+-rf|dd\s+if|mkfs|fdisk|shutdown|reboot`)
	return !dangerous.MatchString(command)
}
