# Course Local AI Code Assistant

Slide: [Download](https://bit.ly/4eUOLmF)

## Before begin

### Required

- Install Ollama -> <https://ollama.com/download>
- Preload model image

  ```bash
  ollama pull llama3.2
  ollama pull codegemma:2b
  ollama pull qwen2.5-coder:1.5b
  ollama pull nomic-embed-text
  ```

### Recommended

- Install Docker -> <https://www.docker.com/products/docker-desktop/>
- Preload docker image

  ```bash
  docker pull ghcr.io/open-webui/open-webui:main
  ```

## Demo

- [Ollama](/1.ollama/README.md)
- [Open WebUI](/2.open-webui/README.md)
- [Continue](/3.continue/README.md)

## Prompt Examples

Dev:

```text
You are senior backend developer that expert in go and postgresql.

[question]

show me step by step.
```

Optimized:

```text
You are a senior sofware developer with expertise in algorithm optimization.
help me to optimization this code:

"""
code snippet
"""
```

SQL:

```text
You are an expert PostgreSQL developer who assists in writing SQL queries.

In PostgreSQL version 16, I have a table defined like this:

"""
Table Schema
"""

[question]

Your response should ONLY include the SQL statement with NO commentary. Make sure the output is valid SQL for PostgreSQL version 16.
```
