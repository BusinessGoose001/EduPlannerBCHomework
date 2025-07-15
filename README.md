# API Developer Homework Assignment 
### Developer: [Gillian Bryson](mailto:gillianfbryson@gmail.com)

## Q1 - List REST endpoints for tables referenced in Q4
### Assumptions
1. We want full CRUD access and no tables are in read-only status.
2. Our App's UI Forms ensure all required data is given in correct format before API is called
3. {{}} denotes a parameter to be filled in by the executing agent
4. On failure to execute the endpoint the appropriate error code will be returned (404, 403, 401, 500 etc.)

### Customer
| Purpose                 | Method | Endpoint                     | Notes                                                                                             |
|-------------------------|--------|------------------------------|---------------------------------------------------------------------------------------------------|
| Create new customer     | POST   | `/customers`                 | Takes JSON form data, Returns customer info object                                                |
| Fetch existing customer | GET    | `/customers/{{customer_id}}` | Returns JSON object of customer info                                                              |
| List customers          | GET    | `/customers`                 | Takes: search parameters to filter results. <br/>Returns List of customer info objects, paginated |
| Update customer         | PATCH  | `/customers/{{customer_id}}` | Takes JSON form data of fields to update.                                                         |
| Delete customer         | DELETE | `/customers/{{customer_id}}` | Returns 204                                                                                       |
- _We would not allow updating of CustomerID_
### Order
| Purpose                  | Method | Endpoint                            | Notes                                                                                                                               | 
|--------------------------|--------|-------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------|
| Create new order         | POST   | `/orders`                           | Takes JSON form data, Returns order info object                                                                                     |
| Fetch existing order     | GET    | `/orders/{{order_id}}`              | Returns JSON object of order info                                                                                                   |
| List all orders          | GET    | `/orders`                           | Takes search params to filter results. <br/>Returns List of order info objects, paginated                                           |
| List orders for customer | GET    | `/customers/{{customer_id}}/orders` | Takes search params to filter results. <br/>Returns List of order info objects where CustomerID column matches provided customer_id |
| Update order             | PATCH  | `/orders/{{order_id}}`              | Takes JSON form data of fields to update.                                                                                           |
| Fulfill order            | PATCH  | `/orders/{{order_id}}/fulfill`      | Returns 204                                                                                                                         |
| Cancel order             | PATCH  | `/orders/{{order_id}}/cancel`       | Returns 204                                                                                                                         |
| Delete order             | DELETE | `/orders/{{order_id}}`              | Returns 204                                                                                                                         |
- _We would not allow updating of OrderID, and, in most cases, CustomerID_
- _How you want to handle deleting of orders is business dependant. In most cases you would likely want to keep it on record but mark it cancelled (w/o a field for it you can set `OrderFulfilledDate` to NULL), but in development cases you may still want the power to completely delete the record._
- _Cancelling and fulfilling orders can be updated through the Update Order endpoint but if there are knock on actions (like notifying the customer via email etc.) it can be nicer to scope them out to specific endpoints_
- _Order is not a great table name as it is part of the protected `ORDER BY` keyword pair for `SQL` 
### LineItem
| Purpose                                          | Method | Endpoint                              | Notes                                                                                         | 
|--------------------------------------------------|--------|---------------------------------------|-----------------------------------------------------------------------------------------------|
| Create new line item                             | POST   | `/lineitems`                          | Takes JSON form data                                                                          |
| Fetch line item                                  | GET    | `/lineitems/{{lineitem_id}}`          | Returns JSON object of line item info                                                         |
| List all line items                              | GET    | `/lineitems`                          | Takes search params to filter results. <br/>Returns List of line item info objects, paginated |
| List line items in order                         | GET    | `/order/{{order_id}}/lineitems`       | Returns List of line item info objects, paginated                                             |
| List line items customer <br/>has ordered before | GET    | `customers/{{customer_id}}/lineitems` | Takes search params to filter results. Returns List of line item info objects, paginated      |
| Update line item                                 | PATCH  | `lineitem/{{lineitem_id}}`            | Takes JSON form data of fields to update.                                                     |
### CustomerAddress
| Purpose                     | Method | Endpoint                                              | Notes                                                                                           |
|-----------------------------|--------|-------------------------------------------------------|-------------------------------------------------------------------------------------------------|
| Create new customer address | POST   | `/customers/{{customer_id}}/addresses`                | Takes JSON form data; associates address with specified customer                                |
| Fetch customer address      | GET    | `/customers/{{customer_id}}/addresses/{{address_id}}` | Returns JSON object of address info                                                             |
| List all customer addresses | GET    | `/addresses`                                          | Takes search params to filter results. Returns List of customer address info objects, paginated |
| List customer addresses     | GET    | `/customers/{{customer_id}}/addresses`                | Returns list of addresses for a given customer                                                  |
| Update customer address     | PATCH  | `/customers/{{customer_id}}/addresses/{{address_id}}` | Takes JSON form data of fields to change                                                        |
| Delete customer address     | DELETE | `/customers/{{customer_id}}/addresses/{{address_id}}` | Returns 204 on success; removes the address record                                              |
- _We would not allow update of CustomerID_
### ReturnItem
| Purpose                   | Method | Endpoint                                            | Notes                                                                                      |
|---------------------------|--------|-----------------------------------------------------|--------------------------------------------------------------------------------------------|
| Create new return item    | POST   | `/orders/{{order_id}}/returns`                      | Takes JSON form data including `line_item_id`, `quantity`, and `amount_returned`           |
| Fetch return item         | GET    | `/orders/{{order_id}}/returns/{{returned_item_id}}` | Returns JSON object of return info                                                         |
| Fetch all returned items  | GET    | `/returns`                                          | Takes search params to filter results. <br/>Returns List of return info objects, paginated |
| List returns for an order | GET    | `/orders/{{order_id}}/returns`                      | Returns list of returned items for the specified order                                     |
| List returns for customer | GET    | `/customers/{{customer_id}}/returns`                | Returns list of returned items for the specified user                                      |
| Update return item        | PATCH  | `/orders/{{order_id}}/returns/{{returned_item_id}}` | Takes JSON form data to update return details (e.g., quantity adjustment or correction)    |
| Delete return item        | DELETE | `/orders/{{order_id}}/returns/{{returned_item_id}}` | Returns 204 on success; typically restricted to admin or correction workflows              |

## Q2 - Using Go return occurrence of digit in series between 2 numbers
#### _NOTE: Go is developer's choice_
function: CountDigitOccurrencesInSeries

file: [digitcounter.go](./digitcount_api/digitcount/digitcounter.go)

tests: [digitcounter_test.go](./digitcount_api/digitcount/digitcounter_test.go)

#### Running Tests
```bash
cd digitcount_api/digitcount
go test
```

More test cases can be added in [testcases.go](./digitcount_api/digitcount/testcases.go)


## Q3 - Host Q2 as a full API
file: [main.go](./digitcount_api/main.go)
### Requirements
- Go 1.24
- gin 1.10.1
### Running
1. Clone Repository
```bash
git clone git@github.com:BusinessGoose001/EduPlannerBCHomework.git
cd EduPlannerBCHomework/digitcount_api
```
2. Start the server
```bash
go run main.go
```
3. Testing
   1. Use Insomnia / Postman to import the cURL command in [test_digit_counter.sh](./digitcount_api/test_digit_counter.sh)
   2. Use Shell to run cURL command [test_digit_counter.sh](./digitcount_api/test_digit_counter.sh)
User is welcome to change file parameters as needed for testing
4. Notes:
   1. To run on a different port, set an environment variable:
      ```bash
      export API_PORT=9090
      ```
      Then run:
      ```bash
      go run main.go
      ```
   2. To properly deploy this to somewhere like AWS or other cloud platforms, we would probably want to containerise the app via docker. This would allow for easier repeatable deployments and (for larger applications) allow usage through kubernetes for distributed load handling.
## Q4 - Queries
### Notes
1. In queries, {{}} denotes a parameter to be filled in by the executing agent (program or user)
_Database backend not specified so assuming we are using a postgres backend._
### Queries
1. [best_selling_product.sql](./q4_queries/best_selling_product.sql)
2. [top_buying_customers.sql](./q4_queries/top_buying_customers.sql)
3. Two parts 
   1. [age_group_sales](./q4_queries/age_groups_sales.sql)
   2. [age_group_country_sales.sql](./q4_queries/age_group_country_sales.sql)
4. [orders_by_day_histogram](./q4_queries/orders_by_day_histogram.sql)
5. [percent_items_discounted](./q4_queries/percent_items_discounted.sql)