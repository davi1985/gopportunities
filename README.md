# 🚀 Gopportunities

**Gopportunities** is a RESTful API built with Go (Golang) designed to manage and publish job openings. This project focuses on simplicity, high performance, and idiomatic Go practices.

---

## 🏗️ Architecture & Project Structure

This project follows a simplified variation of the **Standard Go Project Layout**, ensuring separation of concerns, ease of maintenance, and scalability.

```text
.
├── cmd/
│   └── api/
│       └── main.go           # Application entry point (server/routes initialization)
├── config/                   # Global configuration, env variables, and DB connection
├── internal/                 # Application private code (Business Logic & Delivery)
│   ├── handler/              # HTTP Controllers / REST Handlers
│   ├── repository/           # Persistence layer (Database operations)
│   └── entity/               # Domain entities and models
├── schemas/                  # DTOs, validation schemas, and JSON responses
├── router/                   # Definition and grouping of HTTP routes
├── go.mod
└── go.sum
```

## Tech Stack

- **Language:** [Go](https://go.dev/) (Golang)
- **Web Framework:** [Gin](https://gin-gonic.com/)
- **ORM:** [GORM](https://gorm.io/)
- **Current Database:** SQLite (Development environment)

## Getting Started

### Prerequisites

- **Go** 1.22+ installed
- **Docker** & **Docker Compose** (optional, for containerized setup)

### Running Locally (with SQLite)

1. Clone the repository:

```bash
git clone [https://github.com/davi1985/gopportunities.git](https://github.com/davi1985/gopportunities.git)
cd gopportunities

```

2. Install dependencies:

```bash
go mod download

```

3. Run the application:

```bash
go run main.go

```

_The API will be available at `http://localhost:8080`._

## API Endpoints

| Method   | Endpoint           | Description                      |
| -------- | ------------------ | -------------------------------- |
| `GET`    | `/api/v1/openings` | List all job openings            |
| `GET`    | `/api/v1/opening`  | Get a specific job opening by ID |
| `POST`   | `/api/v1/opening`  | Create a new job opening         |
| `PUT`    | `/api/v1/opening`  | Update an existing job opening   |
| `DELETE` | `/api/v1/opening`  | Delete a job opening             |

## Roadmap & Future Improvements

- [ ] **Standard Go Layout Migration (`cmd/` & `internal/`)**
- Reorganize project structure following the Standard Go Layout.
- Move the entry point to `cmd/api/main.go`.
- Encapsulate handlers, repositories, and domain entities inside the `internal/` directory to protect private packages.

* [ ] **Migration to PostgreSQL / Docker**
* Transition from SQLite (used for quick prototyping) to **PostgreSQL** using Docker Compose.
* Ensure native support for connection pooling and isolated containerized database instances for both testing and production.

* [ ] **Generic Repository Implementation**
* Implement a generic repository pattern using **Go Generics** (`[T any]`) to abstract repetitive CRUD operations across GORM repositories, reducing boilerplate.

```go
type Repository[T any] interface {
    Create(entity *T) error
    FindByID(id uint) (*T, error)
    FindAll() ([]T, error)
    Update(entity *T) error
    Delete(id uint) error
}
```

- [ ] **Automated Test Suite**
- **Unit Testing:** Implement unit tests for handlers, logic, and DTO validations using `testing` and `stretchr/testify`.
- **Integration Testing:** Add support for **Testcontainers** to spin up an actual PostgreSQL container during automated database repository testing.
- [ ] **Full API Containerization**
