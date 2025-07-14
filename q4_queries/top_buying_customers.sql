WITH base_sales AS (
    SELECT
        orders.CustomerID,
        SUM((line.LineItemUnitPrice * line.Quantity) - line.LineItemDiscount) AS CustomerTotalValueSold
    FROM LineItem line
             JOIN "Order" orders ON line.OrderID = orders.OrderID
    GROUP BY orders.CustomerID
),
 base_returns AS (
     SELECT
         orders.CustomerID,
         SUM(returns.AmountReturned) AS CustomerTotalValueReturned
     FROM ReturnItem returns
              JOIN "Order" orders ON returns.OrderID = orders.OrderID
     GROUP BY orders.CustomerID
 )
SELECT
    base_sales.CustomerID,
    base_sales.CustomerTotalValueSold - COALESCE(base_returns.CustomerTotalValueReturned, 0) AS CustomerNetValue
FROM base_sales
         LEFT JOIN base_returns ON base_sales.CustomerID = base_returns.CustomerID
ORDER BY CustomerNetValue DESC
    LIMIT 5;

