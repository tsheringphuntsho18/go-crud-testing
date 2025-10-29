# Go CRUD Testing

A simple RESTful API in Go for managing users, demonstrating CRUD operations and unit testing using the `chi` router.

## Project Structure

The project consists of the following main files:

- `main.go`: Entry point to start the HTTP server.
- `handlers.go`: Contains the core logic for CRUD operations.
- `handlers_test.go`: Contains unit tests for the handlers.
- `go.mod`/`go.sum` - Go module files

```
go-crud-testing/
├── main.go
│
├── handlers.go
│
├── handlers_test.go
|
├── go.mod
|
├── go.sum
│
└── README.md
```

## Features

- Create, Read, Update, and Delete (CRUD) users
- In-memory data storage (no external database required)
- RESTful API endpoints
- Unit tests for handlers

### Prerequisites

- Go 1.18 or higher (recommended)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/tsheringphuntsho18/go-crud-testing.git

   cd go-crud-testing
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

## Running Tests and Analyzing Coverage

To ensure your code works as expected and to visualize which parts are covered by tests, follow these steps:

### 1. Run All Unit Tests

```bash
go test -v
```

The `-v` (verbose) flag shows the name of each test as it runs. All tests should pass.

### 2. Check Code Coverage Percentage

See a summary of your test coverage with:

```bash
go test -v -cover
```

This will print the percentage of code covered by your tests.

### 3. Generate a Visual Coverage Report

Create a detailed, color-coded HTML report showing which lines were tested:

**Step A:** Output the coverage profile to a file:

```bash
go test -coverprofile=coverage.out
```

**Step B:** Convert the profile to an HTML report and open it in your browser:

```bash
go tool cover -html=coverage.out
```

The report will highlight tested code in green and untested code in red, making it easy to spot gaps in your test coverage.
