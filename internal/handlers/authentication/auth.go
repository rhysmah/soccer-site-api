package authentication

type AuthHandler struct {
	userRepo    UserRepository
	authService AuthService
}

// NewAuthHandler creates new authentication handler with dependencies
func NewAuthHandler(userRepo UserRepository, authService AuthService) *AuthHandler {
	return &AuthHandler{
		userRepo:    userRepo,
		authService: authService,
	}
}
