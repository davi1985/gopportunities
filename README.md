# gopportunities

[![Swagger UI](https://img.shields.io/badge/Swagger%20UI-API%20Documentation-85EA2D?logo=swagger&logoColor=black)](https://davi1985.github.io/gopportunities/)

**gopportunities** is a REST API built with Go for managing job openings.

The project uses Gin for HTTP routing, GORM for database access, and SQLite during development.

## Tech Stack

- **Go** 1.22+
- **Gin** — HTTP web framework
- **GORM** — ORM
- **SQLite** — Development database
- **Docker / Docker Compose** — Planned for the database and application environment

## Project Structure

The project currently uses a simple structure with each package focused on a specific responsibility:

```text
.
├── config/                   # Application configuration and database setup
├── handler/                  # HTTP handlers
├── router/                   # HTTP route definitions
├── schemas/                  # Request and response schemas
├── go.mod
├── go.sum
└── main.go                   # Application entry point
```

## Getting Started

### Requirements

- Go 1.22 or later
- Git

Docker is not required to run the current version, since the application uses SQLite.

### Clone the repository

```bash
git clone https://github.com/davi1985/gopportunities.git
cd gopportunities
```

### Install dependencies

```bash
go mod download
```

### Run the API

```bash
go run main.go
```

The API will be available at:

```text
http://localhost:8080
```

## API

The API is available under the `/api/v1` prefix.

| Method   | Endpoint           | Description             |
| -------- | ------------------ | ----------------------- |
| `GET`    | `/api/v1/openings` | List all job openings   |
| `GET`    | `/api/v1/opening`  | Get a job opening by ID |
| `POST`   | `/api/v1/opening`  | Create a job opening    |
| `PUT`    | `/api/v1/opening`  | Update a job opening    |
| `DELETE` | `/api/v1/opening`  | Delete a job opening    |

### API Documentation

The complete API documentation is available through Swagger UI:

**[Open Swagger UI](https://davi1985.github.io/gopportunities/)**

## Roadmap

- [x] Move the application entry point to `cmd/api/main.go`
- [x] Adopt the `internal/` package structure
- [ ] Migrate from SQLite to PostgreSQL
- [ ] Add Docker Compose configuration
- [ ] Add a generic repository using Go Generics
- [ ] Add unit tests for handlers and validations
- [ ] Add integration tests with Testcontainers
- [ ] Containerize the API

### Generic Repository

One of the planned improvements is a generic repository to reduce repeated CRUD code across repositories:

```go
type Repository[T any] interface {
    Create(entity *T) error
    FindByID(id uint) (*T, error)
    FindAll() ([]T, error)
    Update(entity *T) error
    Delete(id uint) error
}
```

### Testing

The test suite will cover different parts of the application:

- **Unit tests:** handlers, application logic, and request validation.
- **Integration tests:** repository and database operations using a PostgreSQL container through Testcontainers.

## License

This project is available for learning and experimentation.
