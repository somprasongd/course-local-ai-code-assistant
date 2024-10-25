# Continue

## Installation

[VS Code Extension](https://marketplace.visualstudio.com/items?itemName=Continue.continue)

## Configuration

Config.json

```json
{
  "models": [
    {
      "apiBase": "http://localhost:11434/",
      "title": "qwen2.5-coder 7b instruct",
      "provider": "ollama",
      "model": "qwen2.5-coder:7b-instruct"
    }
  ],
  "tabAutocompleteModel": {
    "apiBase": "http://localhost:11434/",
    "title": "Codegemma 2B",
    "provider": "ollama",
    "model": "codegemma:2b-code"
  },
  "tabAutocompleteOptions": {
    "disable": false,
    "debounceDelay": 500,
    "maxPromptTokens": 1500,
    "multilineCompletions": "always",
    "disableInFiles": [
      "*.md"
    ]
  },
  "embeddingsProvider": {
    "apiBase": "http://localhost:11434/",
    "provider": "ollama",
    "model": "nomic-embed-text"
  },
  ...
}
```

## Demo

- [Todo API](/3.continue/demo/todo_api.md)
- [/test](/3.continue/demo/unit_test.md)
- [/comment](/3.continue/demo/unit_test.md)
- [/commit](/3.continue/demo/commit.md)
- @Docs
- [@Postgresql](/3.continue/demo/postgresql.md)
- Prompt file
