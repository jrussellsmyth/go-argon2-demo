package main

import (
	"crypto/rand"
	"crypto/subtle"
	"fmt"

	"golang.org/x/crypto/argon2"
)

// The pepper is a global secret that should be retrieved from a secure source.
// In a real application, you'd use os.Getenv() to load it.
var pepper = "secret-pepper"

// Argon2 parameters
const (
	argon2Time    = 1         // Number of iterations
	argon2Memory  = 64 * 1024 // Memory usage in KB
	argon2Threads = 4         // Number of threads
	argon2KeyLen  = 32        // Length of the derived key
	saltLength    = 16        // Length of the salt
)

// generateSalt creates a random salt for password hashing
func generateSalt() ([]byte, error) {
	salt := make([]byte, saltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

// hashPasswordArgon2 creates an Argon2 hash of the password with salt
func hashPasswordArgon2(password string, salt []byte) []byte {
	return argon2.IDKey([]byte(password), salt, argon2Time, argon2Memory, argon2Threads, argon2KeyLen)
}

// --- The following would interact with a database ---

// User represents the data stored in the database.
type User struct {
	ID           int
	Username     string
	PasswordHash []byte
	Salt         []byte
}

var db = make(map[string]User)

// RegisterUser hashes the password (with pepper) and "stores" the new user.
func RegisterUser(username, password string) error {
	// Generate a random salt
	salt, err := generateSalt()
	if err != nil {
		return fmt.Errorf("failed to generate salt: %w", err)
	}

	// Concatenate the password and pepper before hashing.
	passwordWithPepper := password + pepper

	// Hash the password using Argon2
	hashedPassword := hashPasswordArgon2(passwordWithPepper, salt)

	// Simulate storing the user in a database.
	db[username] = User{
		ID:           len(db) + 1,
		Username:     username,
		PasswordHash: hashedPassword,
		Salt:         salt,
	}
	return nil
}

// AuthenticateUser retrieves a user and checks their password (with pepper).
func AuthenticateUser(username, password string) bool {
	// 1. Retrieve the user from the "database".
	user, ok := db[username]
	if !ok {
		fmt.Println("User not found.")
		return false
	}

	// 2. Concatenate the provided password and pepper for comparison.
	passwordWithPepper := password + pepper

	// 3. Hash the provided password with the stored salt using Argon2
	hashedPassword := hashPasswordArgon2(passwordWithPepper, user.Salt)

	// 4. Compare the hashes using constant time comparison to prevent timing attacks
	if subtle.ConstantTimeCompare(user.PasswordHash, hashedPassword) == 1 {
		return true
	}

	fmt.Println("Invalid password.")
	return false
}

func main() {
	// For this example, we'll set the pepper. In production, this should be from a secure source.

	// Simulate user registration
	fmt.Println("Simulating user registration:")
	if err := RegisterUser("john_doe", "my-super-secret-password-123"); err != nil {
		fmt.Println("\tRegistration failed:", err)
	} else {
		fmt.Println("\tRegistration succeeded.")
	}

	// Simulate successful authentication
	fmt.Println("\nSimulating successful authentication:")
	if !AuthenticateUser("john_doe", "my-super-secret-password-123") {
		fmt.Println("\tAuthentication failed.")
	} else {
		fmt.Println("\tAuthentication succeeded.")
	}

	// Simulate failed authentication (wrong password)
	fmt.Println("\nSimulating failed authentication with wrong password:")
	if !AuthenticateUser("john_doe", "wrong-password") {
		fmt.Println("\tAuthentication failed.")
	} else {
		fmt.Println("\tAuthentication succeeded.")
	}

	// Simulate failed authentication (non-existent user)
	fmt.Println("\nSimulating authentication for non-existent user:")
	if !AuthenticateUser("jane_doe", "some-password") {
		fmt.Println("\tAuthentication failed.")
	} else {
		fmt.Println("\tAuthentication succeeded.")
	}
}
