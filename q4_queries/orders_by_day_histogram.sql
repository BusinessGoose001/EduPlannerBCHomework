WITH base AS (
    SELECT
        EXTRACT(DAY FROM AGE(OrderFulfilledDate, OrderCreatedDate)) AS FulfilmentTime
    FROM
        "Order" orders
    WHERE OrderFulfilledDate IS NOT NULL
)
SELECT
    COUNT(*) as OrdersCount,
CASE
    WHEN FulfilmentTime < 1 THEN 'Less than a day'
    WHEN FulfilmentTime BETWEEN 1 AND 2 THEN '1-2'
    WHEN FulfilmentTime BETWEEN 2 AND 3 THEN '2-3'
    WHEN FulfilmentTime BETWEEN 3 AND 4 THEN '3-4'
    ELSE 'Greater than 4 days'
END as FulfilmentWindow
GROUP BY FulfilmentWindow;