# API Developer Homework Assignment 
### Developer: [Gillian Bryson](mailto:gillianfbryson@gmail.com)

## Q1 - List REST endpoints for tables refrenced in Q4
### Assumptions
1. We want full CRUD access and, no tables are in read-only status.
2. Our App's UI Forms ensure all required data is given in correct format before API is called

### Table: Customer
| Purpose                 | Method | Endpoint                 | Notes                                                                                             |
|-------------------------|--------|--------------------------|---------------------------------------------------------------------------------------------------|
| Create New Customer     | POST   | /customers               | Takes JSON form data                                                                              |
| Fetch Existing Customer | GET    | /customers/{customer_id} | Returns 200, JSON object of customer info                                                         |
| List Customers          | GET    | /customers               | Takes: search parameters to filter results. Returns 200, List of customer info objects, paginated |
| Update Customer         | PATCH  | /customers/{customer_id} | Takes JSON form data of fields to change                                                          |
| Delete Customer         | DELETE | /customers/{customer_id} | Returns 204, if successful, else appropriate error code (404, 403, 401 etc)                       |


## Q2 - Using Go return occurrence of digit in series between 2 numbers
#### _NOTE: Go is developer's choice_


## Q3 - Host Q2 as a full API

## Q4 - 
