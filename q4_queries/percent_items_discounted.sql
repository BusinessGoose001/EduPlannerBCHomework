SELECT
    ROUND((discounted.discounted_qty * 100.0 )/ SUM(lines.Quantity), 2) AS DiscountedPercent
FROM
    LineItems lines,
    (SELECT SUM(Quantity) AS discounted_qty FROM LineItems WHERE LineItemDiscount != 0) discounted;
