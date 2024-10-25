# Write SQL

## Prompt

**System:**

```text
Given the following SQL tables, your job is to write queries given a user’s request.
    
    CREATE TABLE Orders (
      OrderID int,
      CustomerID int,
      OrderDate datetime,
      OrderTime varchar(8),
      PRIMARY KEY (OrderID)
    );
    
    CREATE TABLE OrderDetails (
      OrderDetailID int,
      OrderID int,
      ProductID int,
      Quantity int,
      PRIMARY KEY (OrderDetailID)
    );
    
    CREATE TABLE Products (
      ProductID int,
      ProductName varchar(50),
      Category varchar(50),
      UnitPrice decimal(10, 2),
      Stock int,
      PRIMARY KEY (ProductID)
    );
    
    CREATE TABLE Customers (
      CustomerID int,
      FirstName varchar(50),
      LastName varchar(50),
      Email varchar(100),
      Phone varchar(20),
      PRIMARY KEY (CustomerID)
    );
```

**User:**

```text
Write a SQL query which computes the average total order value for all orders on 2023-04-01.
```

## Sample Response

Here's how you can compute the average total order value:

```sql
SELECT AVG(TotalValue) as AverageTotalOrderValue 
FROM (
    SELECT SUM(OrderDetails.Quantity * Products.UnitPrice) as TotalValue 
    FROM Orders 
    JOIN OrderDetails ON Orders.OrderID = OrderDetails.OrderID 
    JOIN Products ON OrderDetails.ProductID = Products.ProductID 
    WHERE Orders.OrderDate = '2023-04-01' 
    GROUP BY OrderID
) subquery;
```

This query works by first creating a subquery that calculates the total value of each order on 2023-04-01. It does this by multiplying the quantity of each product in an order with the unit price of that product, and then summing these values for each order.

The outer query then computes the average of these totals.
