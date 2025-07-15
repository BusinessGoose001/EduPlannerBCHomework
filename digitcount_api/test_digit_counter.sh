curl -X POST http://localhost:8080/findInSeries \
  -H "Content-Type: application/json" \
  -d '{
    "series_start": 0,
    "series_end": 11,
    "series_increment": 4,
    "series_type": 1,
    "specified_digit": 1
  }'