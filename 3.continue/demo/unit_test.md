# /test

Create `main.go`

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

## Prompt

```text
/test @File main.go
```

## Output

```go
package main

import (
 "testing"
)

func TestFibonacci(t *testing.T) {
 testCases := []struct {
  name     string
  n        int
  expected int
 }{
  {"Base case 1", 0, 0},
  {"Base case 2", 1, 1},
  {"Small number", 5, 5},
  {"Large number", 10, 55},
  {"Negative number", -3, 0}, // Assuming negative numbers return 0 for simplicity
 }

 for _, tc := range testCases {
  t.Run(tc.name, func(t *testing.T) {
   result := fibonacci(tc.n)
   if result != tc.expected {
    t.Errorf("fibonacci(%d) = %d; expected %d", tc.n, result, tc.expected)
   }
  })
 }
}
```
