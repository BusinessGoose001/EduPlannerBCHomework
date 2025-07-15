WITH base AS (
    SELECT
        SUM((line.LineItemUnitPrice * line.Quantity) - line.LineItemDiscount) AS CustomerTotalValueSold,
        EXTRACT(YEAR FROM AGE(CURRENT_DATE, cust.DOB)) AS Age,
        cust.Country
    FROM LineItem line
             LEFT JOIN "Order" orders ON line.OrderID = orders.OrderID
             LEFT JOIN Customer cust ON orders.CustomerID = cust.CustomerID
    GROUP BY cust.Country, Age
),
age_buckets AS (
 SELECT
     SUM(CustomerTotalValueSold) AS AgeTotalValueSold,
     CASE
         WHEN Age < 18 THEN 'Under 18'
         WHEN Age BETWEEN 18 AND 29 THEN '18–29'
         WHEN Age BETWEEN 30 AND 45 THEN '30–45'
         WHEN Age BETWEEN 46 AND 65 THEN '46–65'
         ELSE '65+'
         END AS AgeGroup,
     Country
 FROM base
 GROUP BY Country, AgeGroup
)
SELECT DISTINCT ON (AgeGroup)
    AgeGroup,
    Country,
    AgeTotalValueSold
FROM age_buckets
ORDER BY AgeGroup, AgeTotalValueSold DESC;