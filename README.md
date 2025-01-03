# Nigerian Banks API

A simple REST API that provides information about Nigerian banks, including their names, codes, and USSD codes.

## Features

- Get list of Nigerian banks
- SQLite database for data storage
- Rate limiting (5 requests per second with burst of 10)

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
