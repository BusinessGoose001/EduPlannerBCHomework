# API Developer Homework Assignment 
### Developer: [Gillian Bryson](mailto:gillianfbryson@gmail.com)

## Q1 - List REST endpoints for tables refrenced in Q4
### Assumptions
1. We want full CRUD access and, no tables are in read-only status.
2. Our App's UI Forms ensure all required data is given in correct format before API is called
3. {} denotes a parameter to be filled in by the executing agent
4. on failure to execute the endpoint the appropriate error code will be returned (404, 403, 401, 500 etc)

### Customer
| Purpose                 | Method | Endpoint                   | Notes                                                                                             |
|-------------------------|--------|----------------------------|---------------------------------------------------------------------------------------------------|
| Create new customer     | POST   | `/customers`               | Takes JSON form data                                                                              |
| Fetch existing customer | GET    | `/customers/{customer_id}` | Returns 200, JSON object of customer info                                                         |
| List customers          | GET    | `/customers`               | Takes: search parameters to filter results. Returns 200, List of customer info objects, paginated |
| Update customer         | PATCH  | `/customers/{customer_id}` | Takes JSON form data of fields to change                                                          |
| Delete customer         | DELETE | `/customers/{customer_id}` | Returns 204                                                                                       |
- _We would not allow updating of CustomerID_
### Order
| Purpose                  | Method | Endpoint                           | Notes                                                                                                                               | 
|--------------------------|--------|------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------|
| Create new order         | POST   | `/orders`                          | Takes JSON form data                                                                                                                |
| Fetch existing order     | GET    | `/orders/{order_id}`               | Returns 200, JSON object of order info                                                                                              |
| List all orders          | GET    | `/orders`                          | Takes search params to filter results. Returns 200, List of order info objects, paginated                                           |
| List orders for customer | GET    | `/customers/{customer_id}/orders`  | Takes search params to filter results. Returns 200, List of order info objects where CustomerID column matches provided customer_id |
| Update order             | PATCH  | `/orders/{order_id}`               | Takes JSON form data of fields to update.                                                                                           |
| Fulfill order            | PATCH  | `/orders/{order_id}/fulfill`       | Returns 204                                                                                                                         |
| Cancel order             | PATCH  | `/orders/{order_id}/cancel`        | Returns 204                                                                                                                         |
| Delete order             | DELETE | `/orders/{order_id}`               | Returns 204                                                                                                                         |
- _We would not allow updating of OrderID, and, in most cases, CustomerID_
- _How you want to handle deleting of orders is business dependant. In most cases you would likely want to keep it on record but mark it cancelled, but in development cases you may still want the power to completely delete the record._
- _Cancelling and fulfilling orders can be updated through the Update Order endpoint but if there are knock on actions (like notifying the customer via email etc.) it can be nicer to scope them out to specific endpoints_

## Q2 - Using Go return occurrence of digit in series between 2 numbers
#### _NOTE: Go is developer's choice_


## Q3 - Host Q2 as a full API

## Q4 - 
