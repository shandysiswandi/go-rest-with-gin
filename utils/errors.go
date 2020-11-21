package utils

// AppError is ...
type AppError struct {
	Code    string `json:"code"`
	Error   bool   `json:"error"`
	Message string `json:"message"`
}
