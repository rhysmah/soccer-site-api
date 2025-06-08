package auth

type UserRepository interface {
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
