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
| Create new customer     | POST   | `/customers`               | Takes JSON form data, Returns customer info object                                                |
| Fetch existing customer | GET    | `/customers/{customer_id}` | Returns 200, JSON object of customer info                                                         |
| List customers          | GET    | `/customers`               | Takes: search parameters to filter results. Returns 200, List of <br/>customer info objects, paginated |
| Update customer         | PATCH  | `/customers/{customer_id}` | Takes JSON form data of fields to update.                                                         |
| Delete customer         | DELETE | `/customers/{customer_id}` | Returns 204                                                                                       |
- _We would not allow updating of CustomerID_
### Order
| Purpose                  | Method | Endpoint                           | Notes                                                                                                                               | 
|--------------------------|--------|------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------|
| Create new order         | POST   | `/orders`                          | Takes JSON form data, Returns order info object                                                                                     |
| Fetch existing order     | GET    | `/orders/{order_id}`               | Returns 200, JSON object of order info                                                                                              |
| List all orders          | GET    | `/orders`                          | Takes search params to filter results. Returns 200, List of order info objects, paginated                                           |
| List orders for customer | GET    | `/customers/{customer_id}/orders`  | Takes search params to filter results. Returns 200, List of order info objects where <br/>CustomerID column matches provided customer_id |
| Update order             | PATCH  | `/orders/{order_id}`               | Takes JSON form data of fields to update.                                                                                           |
| Fulfill order            | PATCH  | `/orders/{order_id}/fulfill`       | Returns 204                                                                                                                         |
| Cancel order             | PATCH  | `/orders/{order_id}/cancel`        | Returns 204                                                                                                                         |
| Delete order             | DELETE | `/orders/{order_id}`               | Returns 204                                                                                                                         |
- _We would not allow updating of OrderID, and, in most cases, CustomerID_
- _How you want to handle deleting of orders is business dependant. In most cases you would likely want to keep it on record but mark it cancelled (w/o a field for it you can set `OrderFulfilledDate` to NULL), but in development cases you may still want the power to completely delete the record._
- _Cancelling and fulfilling orders can be updated through the Update Order endpoint but if there are knock on actions (like notifying the customer via email etc.) it can be nicer to scope them out to specific endpoints_
### LineItem
| Purpose                                     | Method | Endpoint                           | Notes                                                                                         | 
|---------------------------------------------|--------|------------------------------------|-----------------------------------------------------------------------------------------------|
| Create new line item                        | POST   | `/lineitems/`                      | Takes JSON form data                                                                          |
| Fetch line item                             | GET    | `/lineitems/{lineitem_id}`         | Returns 200, JSON object of line item info                                                    |
| List all line items                         | GET    | `/lineitems/`                      | Takes search params to filter results. Returns 200, List of line item info objects, paginated |
| List line items in order                    | GET    | `/order/{order_id}/lineitems`      | Returns 200, List of line item info objects, paginated                                        |
| List line items customer has ordered before | GET    | `customer/{customer_id}/lineitems` | Takes search params to filter results. Returns 200, List of line item info objects, paginated |
| Update line item                            | PATCH  | `lineitem/{lineitem_id}`           | Takes JSON form data of fields to update.                                                     |
- _To preserve historical accuracy, pricing should ideally be stored in a separate `ItemPrice` table or equivalent. With teh database as provided, updating `LineItemUnitPrice` could alter order history on recall. An alternative is generating static receipts at purchase time and storing them externally which is simpler, but less useful for long-term analytics._

## Q2 - Using Go return occurrence of digit in series between 2 numbers
#### _NOTE: Go is developer's choice_


## Q3 - Host Q2 as a full API

## Q4 - 
