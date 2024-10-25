# Write SQL

## Prompt

```text
Instructions:
Your task is to convert a question into a SQL query, given a Postgres database schema.

Adhere to these rules:

- Deliberately go through the question and database schema word by word to appropriately answer the question
- Use Table Aliases to prevent ambiguity. For example, SELECT table1.col1, table2.col1 FROM table1 JOIN table2 ON table1.id = table2.id.
- When creating a ratio, always cast the numerator as float

Input:
Generate a SQL query that answers the question which products generate the most sales.

This query will run on a database whose schema is represented in this string:

CREATE TABLE products (

product_id INTEGER PRIMARY KEY, -- Unique ID for each product

name VARCHAR(50), -- Name of the product

price DECIMAL(10,2), -- Price of each unit of the product

quantity INTEGER -- Current quantity in stock

);

CREATE TABLE customers (

customer_id INTEGER PRIMARY KEY, -- Unique ID for each customer

name VARCHAR(50), -- Name of the customer

address VARCHAR(100) -- Mailing address of the customer

);

CREATE TABLE salespeople (

salesperson_id INTEGER PRIMARY KEY, -- Unique ID for each salesperson

name VARCHAR(50), -- Name of the salesperson

region VARCHAR(50) -- Geographic sales region

);

CREATE TABLE sales (

sale_id INTEGER PRIMARY KEY, -- Unique ID for each sale

product_id INTEGER, -- ID of product sold

customer_id INTEGER, -- ID of customer who made purchase

salesperson_id INTEGER, -- ID of salesperson who made the sale

sale_date DATE, -- Date the sale occurred

quantity INTEGER -- Quantity of product sold

);

CREATE TABLE product_suppliers (

supplier_id INTEGER PRIMARY KEY, -- Unique ID for each supplier

product_id INTEGER, -- Product ID supplied

supply_price DECIMAL(10,2) -- Unit price charged by supplier

);

-- sales.product_id can be joined with products.product_id

-- sales.customer_id can be joined with customers.customer_id

-- sales.salesperson_id can be joined with salespeople.salesperson_id

-- product_suppliers.product_id can be joined with products.product_id

Response:
Based on your instructions, here is the SQL query I have generated to answer the question which products generate the most sales.:
```

## Sample Response

```sql
SELECT 
    p.product_id,
    p.name,
    SUM(s.quantity) AS total_quantity_sold
FROM 
    sales s
JOIN 
    products p ON s.product_id = p.product_id
GROUP BY 
    p.product_id, p.name
ORDER BY 
    total_quantity_sold DESC;
```

This query joins the `sales` table with the `products` table on the `product_id` column to get the product details for each sale. It then groups the results by product ID and name, and calculates the total quantity sold for each product. Finally, it orders the results in descending order based on the total quantity sold, so that the products generating the most sales appear first.
