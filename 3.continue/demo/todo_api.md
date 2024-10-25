# Todo API

## Prompt

```text
write a todo api with go and fiber framework using haxagonal and database postgresql, show me step by step
```

## Output

 Sure! Let's create a simple TODO API using Go and the Fiber framework. We'll follow the Hexagonal architecture (also known as Ports and Adapters) and use PostgreSQL for the database. Here’s a step-by-step guide to set up this project.

### Step 1: Set Up Your Project

First, create a new directory for your project and navigate into it:

```sh
mkdir todo-api && cd todo-api
```

Initialize a Go module:

```sh
go mod init github.com/yourusername/todo-api
```

### Step 2: Install Dependencies

Install the necessary dependencies:

```sh
go get -u github.com/gofiber/fiber/v2
go get -u github.com/golang-migrate/migrate/v4
go get -u github.com/jmoiron/sqlx
go get -u github.com/lib/pq
```

### Step 3: Create the Project Structure

Create the following directories and files:

```sh
mkdir cmd internal adapters persistence usecases
touch main.go cmd/app.go internal/server.go adapters/todo_adapter.go adapters/postgres.go persistence/todo.go persistence/migrations/init.sql persistence/migrations/001_create_todos_table.sql
```

### Step 4: Define the Migrations

In `persistence/migrations/001_create_todos_table.sql`, create your database table:

```sql
CREATE TABLE todos (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    completed BOOLEAN NOT NULL DEFAULT FALSE
);
```

### Step 5: Create the Main Application File

In `main.go`, set up the Fiber app and initialize the database:

```go
package main

import (
 "log"
 "os"

 "github.com/gofiber/fiber/v2"
 _ "github.com/lib/pq"
 "github.com/jmoiron/sqlx"
 "github.com/joho/godotenv"
)

func main() {
 err := godotenv.Load()
 if err != nil {
  log.Fatal("Error loading .env file")
 }

 db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
 if err != nil {
  log.Fatalln(err)
 }

 app := fiber.New()

 // Initialize the server with the database connection and Fiber app
 server := &Server{
  App:  app,
  DB:   db,
  TodoAdapter: &TodoAdapter{
   DB: db,
  },
 }

 // Run the migration
 _, err = db.Exec(initSchema)
 if err != nil {
  log.Fatalln("Failed to run migrations:", err)
 }

 // Define your routes here
 server.SetupRoutes()

 port := os.Getenv("PORT")
 if port == "" {
  port = "3000"
 }

 err = app.Listen(":" + port)
 if err != nil {
  log.Fatal(err)
 }
}
```

### Step 6: Define the Hexagonal Adapters

In `adapters/todo_adapter.go`, define the adapter for your TODOs:

```go
package adapters

import (
 "github.com/jmoiron/sqlx"
 "github.com/yourusername/todo-api/persistence"
)

type TodoAdapter struct {
 DB *sqlx.DB
}

func (ta *TodoAdapter) GetTodos() ([]persistence.Todo, error) {
 var todos []persistence.Todo
 err := ta.DB.Select(&todos, "SELECT id, title, completed FROM todos")
 return todos, err
}

func (ta *TodoAdapter) AddTodo(todo persistence.Todo) error {
 _, err := ta.DB.Exec("INSERT INTO todos (title, completed) VALUES ($1, $2)", todo.Title, todo.Completed)
 return err
}
```

### Step 7: Define the Persistence Layer

In `persistence/todo.go`, define your TODO model and repository interface:

```go
package persistence

type Todo struct {
 ID        int    `db:"id"`
 Title     string `db:"title"`
 Completed bool   `db:"completed"`
```

### Step 8: Define the Use Cases

In `usecases/todo_usecase.go`, define the use cases for your TODO operations:

```go
package usecases

import (
 "github.com/yourusername/todo-api/persistence"
)

type TodoUseCase struct {
 Repo persistence.TodoRepository
}

func (uc *TodoUseCase) GetTodos() ([]persistence.Todo, error) {
 return uc.Repo.GetTodos()
}

func (uc *TodoUseCase) AddTodo(todo persistence.Todo) error {
 return uc.Repo.AddTodo(todo)
}
```

### Step 9: Define the Repository Interface

In `persistence/todo_repository.go`, define the repository interface:

```go
package persistence

type TodoRepository interface {
 GetTodos() ([]Todo, error)
 AddTodo(Todo) error
}
```

### Step 10: Implement the PostgreSQL Repository

In `persistence/postgres_repository.go`, implement the repository using PostgreSQL:

```go
package persistence

import (
 "github.com/jmoiron/sqlx"
)

type PostgresRepository struct {
 DB *sqlx.DB
}

func (pr *PostgresRepository) GetTodos() ([]Todo, error) {
 var todos []Todo
 err := pr.DB.Select(&todos, "SELECT id, title, completed FROM todos")
 return todos, err
}

func (pr *PostgresRepository) AddTodo(todo Todo) error {
 _, err := pr.DB.Exec("INSERT INTO todos (title, completed) VALUES ($1, $2)", todo.Title, todo.Completed)
 return err
}
```

### Step 11: Define the Server and Routes

In `internal/server.go`, define the server struct and setup routes:

```go
package internal

import (
 "github.com/gofiber/fiber/v2"
 "github.com/jmoiron/sqlx"
 "github.com/yourusername/todo-api/adapters"
 "github.com/yourusername/todo-api/persistence"
 "github.com/yourusername/todo-api/usecases"
)

type Server struct {
 App         *fiber.App
 DB          *sqlx.DB
 TodoAdapter persistence.TodoRepository
}

func (s *Server) SetupRoutes() {
 s.App.Get("/todos", func(c *fiber.Ctx) error {
  todos, err := s.TodoAdapter.GetTodos()
  if err != nil {
   return c.Status(500).JSON(fiber.Map{"error": "Cannot get todos"})
  }
  return c.JSON(todos)
 })

 s.App.Post("/todos", func(c *fiber.Ctx) error {
  var todo persistence.Todo
  if err := c.BodyParser(&todo); err != nil {
   return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
  }
  err := s.TodoAdapter.AddTodo(todo)
  if err != nil {
   return c.Status(500).JSON(fiber.Map{"error": "Cannot add todo"})
  }
  return c.SendStatus(201)
 })
}
```

### Step 12: Run the Application

Finally, run your application using `go run main.go`. You should now have a basic TODO API running on your local server.

### Conclusion

This completes the setup of a simple TODO API using Go, Fiber, Hexagonal architecture, and PostgreSQL. You can expand this project by adding more features, improving error handling, and enhancing security measures.
