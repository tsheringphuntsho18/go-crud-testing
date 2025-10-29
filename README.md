# Go CRUD Testing

A simple RESTful API in Go for managing users, demonstrating CRUD operations and unit testing using the `chi` router.

## Features

- Create, Read, Update, and Delete (CRUD) users
- In-memory data storage (no external database required)
- RESTful API endpoints
- Unit tests for handlers

## Getting Started

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


### API Endpoints

- `GET    /users` - List all users
- `POST   /users` - Create a new user
- `GET    /users/{id}` - Get a user by ID
- `PUT    /users/{id}` - Update a user by ID
- `DELETE /users/{id}` - Delete a user by ID


## Running Tests

To run the unit tests:

```bash
go test -v
```

## Project Structure

- `main.go` - Entry point, sets up routes and starts the server
- `handlers.go` - Handler functions for each endpoint
- `handlers_test.go`- Unit tests for handlers
- `go.mod`/`go.sum` - Go module files

