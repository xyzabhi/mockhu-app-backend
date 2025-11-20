package dtos

// UserInfo represents user information returned in responses
type UserInfo struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
}

// ErrorResponse is the standard error response format
type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

// ErrorDetail contains error information
type ErrorDetail struct {
	Code    string            `json:"code"`
	Message string            `json:"message"`
	Fields  map[string]string `json:"fields,omitempty"`
}

// Common error codes
const (
	ErrInvalidInput        = "INVALID_INPUT"
	ErrEmailAlreadyExists  = "EMAIL_ALREADY_EXISTS"
	ErrPhoneAlreadyExists  = "PHONE_ALREADY_EXISTS"
	ErrUsernameTaken       = "USERNAME_TAKEN"
	ErrWeakPassword        = "WEAK_PASSWORD"
	ErrOTPInvalid          = "OTP_INVALID"
	ErrOTPExpired          = "OTP_EXPIRED"
	ErrTooManyOTPAttempts  = "TOO_MANY_OTP_ATTEMPTS"
	ErrRefreshTokenInvalid = "REFRESH_TOKEN_INVALID"
	ErrRefreshTokenReused  = "REFRESH_TOKEN_REUSED"
	ErrUnauthorized        = "UNAUTHORIZED"
	ErrForbidden           = "FORBIDDEN"
	ErrNotFound            = "NOT_FOUND"
	ErrInternalServerError = "INTERNAL_SERVER_ERROR"
)
