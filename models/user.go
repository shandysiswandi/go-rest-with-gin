package models

// User is ..
type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// TableName is
func (b *User) TableName() string {
	return "users"
}
