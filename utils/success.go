package utils

// AppSuccess is ...
type AppSuccess struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
