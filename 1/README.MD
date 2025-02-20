# How to run project
1. Create `conf.yml` from `conf.yml.example`, define the right DB connection
2. Run `make migrate-fresh` or `make migrate` for table migration, `migrate-fresh` will do a clean setup (reset all the db)
3. Run `make seed` this will seed starting data for `users` and `products`

### Testing
Import the `Nexmedis.postman_collection.json` and `Nexmedis.postman_environment.json`, There is already a setup list of API that can be used for testing.

Testing user from seeder:
```
email: email@gmail.com
password: 123456
```

## Summary
You are tasked with designing an API for an e-commerce platform. The system must support the following features:
User registration and authentication
Viewing and searching products
Adding items to a shopping cart
Completing a purchase
Design the RESTful endpoints for the above features. Describe your choice of HTTP methods (GET, POST, PUT, DELETE), URL structure, and the expected response formats. Assume that users need to authenticate before performing certain actions (e.g., adding items to the cart).

-----
Answer: this is a simple example API routing for the project description above.

API List:
-  Non-Authenticated
     - `/auth/register[POST]`  => Return JWT token (request: email, password,name)
     - `/auth/login[POST]` => Return JWT Token (request: email, password)
     - `/products/filter[POST]` =>  filtering all products with `search` and `pagination` 
     - `/products/{id}[GET]` => get product by id
-  Authenticated
      - `/carts/filter[POST]` => Return list of carts that only own by current logged in user (check `user_id` == `current_user.id`)
      - `/carts/{id}[GET]` => return cart if it exist, and the cart belongs to the current user
      - `/carts [POST]` => create new cart
      - `/carts/{id} [PUT]` => Update Qty of the cart, this part can use `PATCH` as well, since it only change the `qty` of the cart, cart must be own by current logged in user
      - `/carts/{id} [DELETE]` => Delete the cart, cart must be own by current logged in user
      - `/transactions/checkout [POST]` => checkout all carts own by current logged in user, after finished checkout, delete cart
      - `/transactions/filter [POST]` => list all transactions own by current user
      - `/transactions/{id} [GET]` => get transaction by id that own by current user
