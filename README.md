# User Service (Go)

A simple **CRUD backend service** written in Go using Gin.  
Demonstrates **clean architecture, RESTful APIs, and Go best practices**.

---

## Features

- Create, Read, Update, Delete users
- Unique email enforcement
- Thread-safe in-memory storage
- Consistent JSON responses

---

## Tech Stack

- Go (1.21+)
- Gin Web Framework
- UUID (`github.com/google/uuid`)

---

## Running the Service

```bash
go run cmd/user_service/main.go
```

## Running the Service

Runs on: [http://localhost:8080](http://localhost:8080)

---

## Running Tests

Run all service and handler layer tests:

```bash
go test ./...
```

## API Endpoints

| Method | Endpoint   | Description       |
| ------ | ---------- | ----------------- |
| POST   | /users     | Create a new user |
| GET    | /users     | Get all users     |
| GET    | /users/:id | Get user by ID    |
| PUT    | /users/:id | Update a user     |
| DELETE | /users/:id | Delete a user     |

---

## Example Usage (PowerShell)

**Create user:**

```powershell
Invoke-RestMethod -Method Post -Uri "http://localhost:8080/users" `
    -ContentType "application/json" `
    -Body '{"name":"Mathew","email":"mathew@example.com"}'
```

**Get all Users:**

```
Invoke-RestMethod -Method Get -Uri "http://localhost:8080/users"
```

**Get User by id:**

```
Invoke-RestMethod -Method Get -Uri "http://localhost:8080/users/<user-id>"
```

**Delete User by id:**

```
Invoke-RestMethod -Method Delete -Uri "http://localhost:8080/users/<user-id>"

```

**Update user by id :**

```
Invoke-RestMethod -Method Put -Uri "http://localhost:8080/users/<user-id>" `
    -ContentType "application/json" `
    -Body '{"name":"Mathew Mbugua","email":"mathew.mbugua@example.com"}'
```

## Notes

- In-memory storage — all data is lost on server restart.
- Can be extended to MongoDB or PostgreSQL without changing handlers.

### Server Features

- Uses **graceful shutdown** to handle system interrupts.  
  The server will finish ongoing requests before stopping, ensuring no abrupt termination of in-progress API calls
