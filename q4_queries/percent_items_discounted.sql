SELECT
    ROUND(discounted.discounted_qty  / SUM(lines.Quantity)* 100.0, 2) AS DiscountedPercent
FROM
    LineItems lines,
    (SELECT SUM(Quantity) AS discounted_qty FROM LineItems WHERE LineItemDiscount != 0) discounted;
