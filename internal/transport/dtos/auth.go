package dtos

// POST /v1/auth/signup
type SignupRequest struct {
	TempID      string `json:"temp_id,omitempty"`
	Method      string `json:"method" binding:"required"` // "email", "mobile", "google"
	Email       string `json:"email,omitempty"`
	Phone       string `json:"phone,omitempty"`
	CountryCode string `json:"country_code,omitempty"` // e.g., "+91"
	Password    string `json:"password,omitempty"`
	SocialToken string `json:"social_token,omitempty"` // For social login
}

type SignupResponse struct {
	UserID              string `json:"user_id"`
	VerificationNeeded  bool   `json:"verification_needed"`
	VerificationChannel string `json:"verification_channel,omitempty"` // "email" or "mobile"
}

// POST /v1/auth/verify
type VerifyRequest struct {
	UserID string `json:"user_id" binding:"required"`
	Method string `json:"method" binding:"required"` // "email" or "mobile"
	Code   string `json:"code" binding:"required"`   // OTP or verification code
}

type VerifyResponse struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresIn    int       `json:"expires_in"` // seconds
	User         *UserInfo `json:"user"`
}

// POST /v1/auth/login
type LoginRequest struct {
	Identifier string `json:"identifier" binding:"required"` // email or phone
	Password   string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"` // seconds
}

// POST /v1/auth/refresh
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type RefreshResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"` // seconds
}

// POST /v1/auth/logout (Auth Required)
type LogoutRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type LogoutResponse struct {
	Message string `json:"message"`
}

// POST /v1/auth/resend
type ResendRequest struct {
	UserID string `json:"user_id" binding:"required"`
	Method string `json:"method" binding:"required"` // "email" or "mobile"
}

type ResendResponse struct {
	Message string `json:"message"`
}
