# GOLANG Project 

# Install Dependencies

**For Backend** - `go run main.go`

# Routes

**POST METHOD** - `http://localhost:9000/v1/api/data`
{
  "open":" ",
  "high":" ",
  "low":" ",
  "close":" ",
  "volume":" "
}

**GET METHOD** - `http://localhost:9000/v1/api/get-all-data`
    {
        "_id": "628bf53c7832dd81337ac85e",
        "Time": "1970-01-01T05:00:00+05:00",
        "symbol": "AAPL",
        "open": "100.00",
        "high": "120.00",
        "low": "150.00",
        "close": "121.00",
        "volume": "10300"
    }