# Nigerian Banks API

A simple REST API that provides information about Nigerian banks, including their names, codes, and USSD codes.

## Features

- Get list of Nigerian banks
- SQLite database for data storage
- Rate limiting (5 requests per second with burst of 10)
- Clean architecture with controllers and models

## Requirements

- Go 1.21 or higher
- SQLite3

## Installation

1. Clone the repository
2. Install dependencies:
```bash
go mod tidy
```

## Running the API

```bash
go run main.go
```

The server will start on port 8080.

## API Endpoints

### Get All Banks
```
GET /api/banks
```

Response format:
```json
{
  "banks": [
    {
      "id": 1,
      "name": "Access Bank",
      "code": "044",
      "ussd_code": "*901#",
      "base_ussd_code": "901",
      "bank_category": "commercial",
      "internet_banking": true
    },
    ...
  ]
}
```

### Add New Bank
```
POST /api/banks
```

Request body:
```json
{
  "name": "New Bank",
  "code": "123",
  "ussd_code": "*123#",
  "base_ussd_code": "123",
  "bank_category": "commercial",
  "internet_banking": true
}
```

Response: The created bank object with its assigned ID.

## Rate Limiting

The API is rate-limited to 5 requests per second with a burst capacity of 10 requests.
