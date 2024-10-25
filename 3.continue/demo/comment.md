# /comment

Open `main.go`

```go
package main

import "fmt"

func main() {
 fmt.Println(fibonacci(10))
}

func fibonacci(n int) int {
 if n <= 1 {
  return n
 }
 return fibonacci(n-1) + fibonacci(n-2)
}

```

Highlight the `fibonacci` function

```go
func fibonacci(n int) int {
 if n <= 1 {
  return n
 }
 return fibonacci(n-1) + fibonacci(n-2)
}
```

Then press cmd/ctrl+shift+L to add it to chat

```text
/comment
```

## Output

```go
package main

import "fmt"

func main() {
 fmt.Println(fibonacci(10))
}

func fibonacci(n int) int {
 // Base case: if n is less than or equal to 1, return n
 if n <= 1 {
  return n
 }
 // Recursive case: return the sum of the two preceding Fibonacci numbers
 return fibonacci(n-1) + fibonacci(n-2)
}

```
