# CRUD API

## Prompt

**System:**

```text
You are senior backend developer that expert in go and postgresql.
```

**User:**

```text
Write a CRUD API for product from:

- table named product with 3 columns id [uuid], name [string], price [number]
- using fiber framework for http server
- using repository pattern with raw sql and 'sqlx' library

show me step by step
```

## Sample Response

Sure! Below is a step-by-step guide to creating a CRUD API for a `product` table using the Fiber framework, PostgreSQL, and the SQLx library. We'll follow these steps:

1. **Set up your project environment**
2. **Install necessary dependencies**
3. **Create database models**
4. **Implement repository pattern**
5. **Create service layer**
6. **Create controller layer**
7. **Initialize Fiber server and routes**

Let's go through each step in detail.

### Step 1: Set up your project environment

First, create a new directory for your project and initialize it with Go modules:

```bash
mkdir product-crud-api
cd product-crud-api
go mod init product-crud-api
```

### Step 2: Install necessary dependencies

You will need to install the following packages:

- `github.com/gofiber/fiber/v2`
- `github.com/jmoiron/sqlx`
- `github.com/google/uuid`
- PostgreSQL driver for Go, e.g., `lib/pq` (if using Docker, PostgreSQL should be accessible).

Add these dependencies to your `go.mod` file:

```bash
go get github.com/gofiber/fiber/v2
go get github.com/jmoiron/sqlx
go get github.com/google/uuid
go get github.com/lib/pq
```

### Step 3: Create database models

Define a Go struct that represents the `product` table in your PostgreSQL database.

```go
// product.go
package main

import (
    "github.com/google/uuid"
)

type Product struct {
    ID    uuid.UUID `db:"id"`
    Name  string    `db:"name"`
    Price float64   `db:"price"`
}
```

### Step 4: Implement repository pattern

Create a repository to handle all database operations.

```go
// repository.go
package main

import (
    "database/sql"
    "fmt"
    "github.com/google/uuid"
)

type ProductRepository struct {
    DB *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
    return &ProductRepository{DB: db}
}

func (r *ProductRepository) Create(product *Product) error {
    query := `INSERT INTO product (id, name, price) VALUES ($1, $2, $3)`
    _, err := r.DB.Exec(query, product.ID, product.Name, product.Price)
    return err
}

func (r *ProductRepository) Get(id uuid.UUID) (*Product, error) {
    query := `SELECT id, name, price FROM product WHERE id = $1`
    var product Product
    err := r.DB.Get(&product, query, id)
    if err != nil {
        return nil, err
    }
    return &product, nil
}

func (r *ProductRepository) GetAll() ([]*Product, error) {
    query := `SELECT id, name, price FROM product`
    var products []*Product
    err := r.DB.Select(&products, query)
    if err != nil {
        return nil, err
    }
    return products, nil
}

func (r *ProductRepository) Update(id uuid.UUID, updatedProduct *Product) error {
    query := `UPDATE product SET name = $1, price = $2 WHERE id = $3`
    _, err := r.DB.Exec(query, updatedProduct.Name, updatedProduct.Price, id)
    return err
}

func (r *ProductRepository) Delete(id uuid.UUID) error {
    query := `DELETE FROM product WHERE id = $1`
    _, err := r.DB.Exec(query, id)
    return err
}
```

### Step 5: Create service layer

Create a service layer to handle business logic.

```go
// service.go
package main

import (
    "fmt"
)

type ProductService struct {
    Repo *ProductRepository
}

func NewProductService(repo *ProductRepository) *ProductService {
    return &ProductService{Repo: repo}
}

func (s *ProductService) Create(product *Product) error {
    return s.Repo.Create(product)
}

func (s *ProductService) Get(id uuid.UUID) (*Product, error) {
    return s.Repo.Get(id)
}

func (s *ProductService) GetAll() ([]*Product, error) {
    return s.Repo.GetAll()
}

func (s *ProductService) Update(id uuid.UUID, updatedProduct *Product) error {
    return s.Repo.Update(id, updatedProduct)
}

func (s *ProductService) Delete(id uuid.UUID) error {
    return s.Repo.Delete(id)
}
```

### Step 6: Create controller layer

Create a controller layer to handle HTTP requests.

```go
// controller.go
package main

import (
    "net/http"
    "github.com/gofiber/fiber/v2"
    "github.com/google/uuid"
)

type ProductController struct {
    Service *ProductService
}

func NewProductController(service *ProductService) *ProductController {
    return &ProductController{Service: service}
}

func (c *ProductController) Create(c *fiber.Ctx) error {
    var product Product
    if err := c.BodyParser(&product); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
    }
    product.ID = uuid.New()
    if err := c.Service.Create(&product); err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create product"})
    }
    return c.Status(http.StatusCreated).JSON(product)
}

func (c *ProductController) Get(c *fiber.Ctx) error {
    id, err := uuid.Parse(c.Params("id"))
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
    }
    product, err := c.Service.Get(id)
    if err != nil {
        if err == sql.ErrNoRows {
            return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
        }
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get product"})
    }
    return c.JSON(product)
}

func (c *ProductController) GetAll(c *fiber.Ctx) error {
    products, err := c.Service.GetAll()
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get all products"})
    }
    return c.JSON(products)
}

func (c *ProductController) Update(c *fiber.Ctx) error {
    id, err := uuid.Parse(c.Params("id"))
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
    }
    var updatedProduct Product
    if err := c.BodyParser(&updatedProduct); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
    }
    if err := c.Service.Update(id, &updatedProduct); err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update product"})
    }
    return c.JSON(updatedProduct)
}

func (c *ProductController) Delete(c *fiber.Ctx) error {
    id, err := uuid.Parse(c.Params("id"))
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
    }
    if err := c.Service.Delete(id); err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete product"})
    }
    return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Product deleted successfully"})
}
```

### Step 7: Initialize Fiber server and routes

Finally, initialize your Fiber server and define the routes.

```go
// main.go
package main

import (
    "database/sql"
    "fmt"
    "github.com/gofiber/fiber/v2"
    _ "github.com/lib/pq"
    "product-crud-api/repository"
    "product-crud-api/service"
    "product-crud-api/controller"
)

func main() {
    // Connect to PostgreSQL database
    db, err := sqlx.Connect("postgres", "user=your_user dbname=your_db password=your_password sslmode=disable")
    if err != nil {
        panic(err)
    }

    // Initialize repository, service, and controller
    repo := repository.NewProductRepository(db)
    service := service.NewProductService(repo)
    ctrl := controller.NewProductController(service)

    // Create Fiber app
    app := fiber.New()

    // Define routes
    app.Post("/products", ctrl.Create)
    app.Get("/products/:id", ctrl.Get)
    app.Get("/products", ctrl.GetAll)
    app.Put("/products/:id", ctrl.Update)
    app.Delete("/products/:id", ctrl.Delete)

    // Start server
    fmt.Println("Server is running on :8080")
    if err := app.Listen(":8080"); err != nil {
        panic(err)
    }
}
```

### Running the application

Make sure your PostgreSQL server is running, and then run your Go application:

```bash
go run main.go
```

You can now interact with your CRUD API using HTTP requests. For example, to create a new product, you can use `curl` or Postman:

```bash
curl -X POST "http://localhost:8080/products" -H "Content-Type: application/json" -d '{"name": "Example Product", "price": 19.99}'
```

This will create a new product and return its details in JSON format.

That's it! You have now implemented a simple CRUD API using Go, Fiber, and PostgreSQL.
