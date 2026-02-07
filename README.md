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


**Get all Users:**

Invoke-RestMethod -Method Get -Uri "http://localhost:8080/users"
```

**Get all Users:**

Invoke-RestMethod -Method Get -Uri "http://localhost:8080/users"

```

```
