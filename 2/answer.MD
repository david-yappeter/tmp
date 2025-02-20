Users
- Id (primary key)
- username
- email
- created_at

`The default index that is commonly used is B-Tree indexing which used to boost the query for common queries that involve equality, range, or sorting`

So in this case we will used `B-Tree` indexing for all case.

1) Fetch a user by username

Ans: By using `CREATE INDEX on users (username)`, the `username` column will be indexed and improve the query. Note that the query for pattern matching using `ILIKE %pattern%`, If the `%` is applied on both left and right of the pattern, the indexing will be useless, the way to make the index still applied while using pattern matching is to only use `%` on either left or right side, for example `ILIKE pattern%`, this will match `username` that start with `pattern`

2) Fetch users who signed up after a certain date (created_at > "2023-01-01")

Ans: Using the index `CREATE INDEX on users (created-at)` the `created_at` column will be indexed and improve the query, this index also support for `range` query, that fit for this case, which is `SELECT * FROM users WHERE created_at > '2023-01-01'` that used `>` greater than query.

3) Fetch a user by email

Ans: By using `CREATE INDEX on users (email)`, the `email` column will be indexed and improve the query. Note that the query for pattern matching using `ILIKE %pattern%`, If the `%` is applied on both left and right of the pattern, the indexing will be useless, the way to make the index still applied while using pattern matching is to only use `%` on either left or right side, for example `ILIKE pattern%`, this will match `email` that start with `pattern`

Since all indexing use the same index type `B-Tree (default index)`, the trade-off is

Read: the query read will be faster for the query that used the indexed key <br/>
Write: for each mutation (insert, update, delete), the database need to recalculate the index again. This must be done so the query readtime can be consistent because of the indexing
