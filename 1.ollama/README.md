# Ollama

## Installation

Download: [https://ollama.com/download](https://ollama.com/download)

Docker: [https://hub.docker.com/r/ollama/ollama](https://hub.docker.com/r/ollama/ollama)

## Demo

Find and fix bugs in source code.

### CLI

**Run model:**

```bash
ollama run llama3.1
>>> Send a message (/? for help)
```

**Prompt:**

```bash
You will be provided with a piece of Python code, and your task is to find and fix bugs in it.
import Random
    a = random.randint(1,12)
    b = random.randint(1,12)
    for i in range(10):
        question = "What is "+a+" x "+b+"? "
        answer = input(question)
        if answer = a*b
            print (Well done!)
        else:
            print("No.")
```

**Sample response:**

There are several bugs in the provided code:

1. `import Random` should be `import random`. In Python, the module name
is case-sensitive.

2. The variables `a` and `b` are strings when concatenated with the
question string, but integers when compared to the user's input. They need
to be converted to strings before concatenation and back to integers for
comparison. However, it would be more idiomatic to store them as strings
from the start.

3. The line `if answer = a*b:` should be `if int(answer) == a*b:` to
convert the user's input into an integer for comparison with `a*b`. Also,
`=` is assignment operator, not comparison operator. It should be `==`.

4. There's no check if the user's input is numeric (either an integer or a
float). If they enter something non-numeric, it will throw an error.

5. The variable `answer` is not defined before being used in `int(answer)`
and the string concatenation with `question`.

6. The last line of the question string has an extra space at the end
which could lead to issues if you're planning to use this code in a
real-world application where UI/UX matters.

Here's how the corrected code should look like:

```python
import random

for _ in range(10):
    a = str(random.randint(1,12))
    b = str(random.randint(1,12))
    question = "What is "+a+" x "+b+"? "
    answer = input(question)

    try:
        if int(answer) == int(a)*int(b):
            print("Well done!")
        else:
            print("No.")
    except ValueError:
        print("Invalid input. Please enter a number.")
```

This code stores `a` and `b` as strings, converts the user's input to an
integer for comparison, and catches any `ValueError` that might occur if
the user enters something non-numeric.

### API

**Doc:** <https://github.com/ollama/ollama/blob/main/docs/api.md>

```bash
curl http://localhost:11434/api/generate -d '{
  "model": "llama3.1",
  "stream": false,
  "prompt": "You will be provided with a piece of Python code, and your task is to find and fix bugs in it.\nimport Random\n    a = random.randint(1,12)\n    b = random.randint(1,12)\n    for i in range(10):\n        question = \"What is \"+a+\" x \"+b+\"? \"\n        answer = input(question)\n        if answer = a*b\n            print (Well done!)\n        else:\n            print(\"No.\")"
}'
```

**Using system:**

```bash
curl http://localhost:11434/api/generate -d '{
  "model": "llama3.1",
  "stream": false,
  "system": "You will be provided with a piece of Python code, and your task is to find and fix bugs in it.",
  "prompt": "import Random\n    a = random.randint(1,12)\n    b = random.randint(1,12)\n    for i in range(10):\n        question = \"What is \"+a+\" x \"+b+\"? \"\n        answer = input(question)\n        if answer = a*b\n            print (Well done!)\n        else:\n            print(\"No.\")"
}'
```

**Sample response:**

```json
{
  "model": "llama3.1",
  "created_at": "2024-10-25T05:07:43.045263Z",
  "response": "There are several bugs in the provided code. Here's the corrected version:\n\n```python\nimport random  # Python module is named \"random\", not \"Random\"\n\n# Generate two random integers between 1 and 12\na = random.randint(1, 12)\nb = random.randint(1, 12)\n\nfor i in range(10):\n    question = f\"What is {a} x {b}?\"\n    answer = input(question)\n    \n    # Corrected comparison operator from \"==\" to \"!=\"\n    if int(answer) != a * b:\n        print(\"No.\")\n    else:\n        print(\"Well done!\")\n```\n\nExplanation of changes:\n\n1. The correct import statement for the `random` module is `import random`, not `import Random`.\n2. In Python, it's good practice to use an f-string (formatted string) to insert variables into strings, making the code more readable.\n3. When comparing user input with a calculated value, you should compare them after converting both values to the same data type (in this case, integers). The `int()` function is used for this purpose.\n4. The correct comparison operator in this context would be \"not equal\" (`!=`), not \"equal\" (`==`). If the user's answer is not equal to the calculated value, print \"No.\".",
  "done": true,
  "done_reason": "stop",
  "context": [],
  "total_duration": 14889352625,
  "load_duration": 26730125,
  "prompt_eval_count": 111,
  "prompt_eval_duration": 186359000,
  "eval_count": 274,
  "eval_duration": 14675226000
}
```
