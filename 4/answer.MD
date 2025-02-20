orders
- id int primary key
- customer_id int (INDEXED)
- product_id int
- order_date timestamp
- amount decimal(10,2)


1) Write an optimized SQL query to find the top 5 customers who spent the most money in the past month.

Answer: `SELECT customer_id, SUM(amount) as total_amount * FROM orders WHERE order_date >= NOW() - INTERVAL '1 month' GROUP BY customer_id ORDER BY total_amount DESC LIMIT 5`

This will query out `customer_id` and `total_amount` that has been spent by the customer by the last month (from now to 1 month ago) and LIMIT by top 5 spenders. Note that we can `INNER JOIN` with `customers` table to the customer name if needed.


2) How would you improve the performance of this query in a production environment?

`The default index that is commonly used is B-Tree indexing which used to boost the query for common queries that involve equality, range, or sorting`

the column that we need to index is `order_date` since we use WHERE `>=` in the `order_date` field.

We didn't need to add indexing for `amount` since it only used for `SUM(amount)`, but if the query used other where query, or sorting from `amount`, the field `amount` will be good to be indexed.
