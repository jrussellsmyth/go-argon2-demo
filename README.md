# Demonstrating Argon2 in Go

A simple Go CLI application demonstrating secure password hashing with Argon2.

## Features

- Secure password hashing using Argon2
- User registration and authentication simulation
- In-memory user storage (for demonstration purposes)
- Pepper-based password strengthening

## Usage

### Run the application
```bash
go run main.go
```

### Build the application
```bash
go build -o argon-demo main.go
```

### Run the compiled binary
```bash
./argon-demo
```

## Security Features

- **Argon2 Password Hashing**: Uses the modern Argon2 algorithm, winner of the Password Hashing Competition
- **Random Salt Generation**: Each password gets a unique cryptographically secure salt
- **Pepper Support**: Additional global secret for enhanced security
- **Constant-Time Comparison**: Prevents timing attack vulnerabilities
- **Configurable Parameters**: Tunable memory, time, and thread parameters

## Configuration

The Argon2 parameters can be adjusted in the constants section:

```go
const (
    argon2Time    = 1      // Number of iterations
    argon2Memory  = 64*1024 // Memory usage in KB
    argon2Threads = 4      // Number of threads
    argon2KeyLen  = 32     // Length of the derived key
    saltLength    = 16     // Length of the salt
)
```

## Dependencies

- Go 1.21+
- `golang.org/x/crypto/argon2`

## Installation

1. Clone the repository
2. Run `go mod tidy` to install dependencies
3. Build or run the application

## Example Output

```
Simulating user registration:
    Registration succeeded.

Simulating successful authentication:
    Authentication succeeded.

Simulating failed authentication with wrong password:
Invalid password.
    Authentication failed.

Simulating authentication for non-existent user:
User not found.
    Authentication failed.
```

## Note

This is a demonstration application. In a production environment:
- Use a proper database instead of in-memory storage
- Load the pepper from environment variables or secure configuration
- Implement proper error handling and logging
- Add input validation and sanitization