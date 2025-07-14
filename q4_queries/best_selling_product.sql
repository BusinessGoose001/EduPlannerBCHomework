WITH base_data AS (
    SELECT
        items.ProductID,
        SUM((items.LineItemUnitPrice * items.Quantity) - items.LineItemDiscount) AS ProductTotalValueSold,
        SUM(items.Quantity) AS ProductTotalUnitsSold
    FROM LineItems items
             LEFT JOIN "Order" orders ON items.OrderID = orders.OrderID
    WHERE orders.OrderCreatedDate BETWEEN DATE({{ start_date }}) AND DATE({{ end_date }})
GROUP BY items.ProductID
    ),
    revenue_rank AS (
SELECT 'Most Revenue' AS RankType, *
FROM base_data
ORDER BY ProductTotalValueSold DESC
    LIMIT 1
    ),
    units_rank AS (
SELECT 'Most Units' AS RankType, *
FROM base_data
ORDER BY ProductTotalUnitsSold DESC
    LIMIT 1
    )
SELECT * FROM revenue_rank
UNION
SELECT * FROM units_rank;