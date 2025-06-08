package authentication

import "time"

// UserRepository defines how user data is stored and retrieved
// Database operations return only errors
type UserRepository interface {
	CreateUser() error
	FindUserByID() (*User, error)
	FindUserByEmail() (*User, error)
	UpdateUser() error
}

type AuthService interface {
	// HashPassword converts a plain-text password into a secure hash
	HashPassword(password string) ([]byte, error)

	// VerifyPassword checks if a plain-text password matches a stored hash
	VerifyPassword(password string, hash []byte) error

	// GenerateToken creates a JWT for authenticated users
	GenerateToken(userID, email string) (string, error)

	// ValidateToken checks if a JWT token is valid and extracts user info
	ValidateToken(token string) (string, error)
}

type User struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
