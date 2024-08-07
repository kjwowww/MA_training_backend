package model

// SignInRequest represents a request for signing in
type SignInRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// TokenResponse represents a response containing a JWT token
type TokenResponse struct {
	Type        string `json:"type"`
	AccessToken string `json:"access_token"`
}

// SignUpRequest represents a request for signing up
type SignUpRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
}
