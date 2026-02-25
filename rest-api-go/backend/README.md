# REST API Go Documentation

## Overview

This project is a simple REST API written in Go, structured with a clean architecture approach.

## Project Structure

- `cmd/main.go`: Entry point for the application.
- `internal/handler/`: HTTP handlers (e.g., `user_handler.go`).
- `internal/model/`: Data models (e.g., `user.go`).
- `internal/repository/`: Data access layer (e.g., `user_repository_memory.go`, `user_respository.go`).
- `internal/service/`: Business logic layer (e.g., `user_service.go`).

## How to Run

1. Install Go (>=1.18).
2. Clone the repository:
   ```bash
   git clone https://github.com/Dhiooonn/golang-docs.git
   cd rest-api-go
   ```
3. Download dependencies:
   ```bash
   go mod tidy
   ```
4. Run the application:
   ```bash
   go run cmd/main.go
   ```

## API Endpoints

- `/users` (GET): List all users
- `/users` (POST): Create a new user
- `/users/{id}` (GET): Get user by ID
- `/users/{id}` (PUT): Update user by ID
- `/users/{id}` (DELETE): Delete user by ID

## Dependencies

- `golang.org/x/crypto/bcrypt`: For password hashing

## Contribution

1. Fork the repository
2. Create a new branch
3. Commit your changes
4. Push to your branch
5. Create a pull request

## License

MIT
