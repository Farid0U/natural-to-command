# AI CLI Translator (Go)

A Go CLI tool that translates human language commands to Linux commands using OpenRouter AI models.

## Features

- Natural language to command translation
- Safety validation
- Interactive mode
- User confirmation before execution

## Build

```bash
go build -o ai-cli-go
```

## Usage

Set the OPENROUTER_API_KEY environment variable:

**Linux/macOS:**
```bash
export OPENROUTER_API_KEY=your_api_key_here
```

**Windows Command Prompt:**
```cmd
set OPENROUTER_API_KEY=your_api_key_here
```

**Windows PowerShell:**
```powershell
$env:OPENROUTER_API_KEY="your_api_key_here"
```

Then run:

**Command Mode:**
```bash
./ai-cli-go -command "list files in current directory"
```

**Interactive Mode:**
```bash
./ai-cli-go -interactive
```

**With Custom Model:**
```bash
./ai-cli-go -command "show disk usage" -model "anthropic/claude-3-haiku"
```

## Safety

The tool validates commands to prevent destructive operations and requires user confirmation.
