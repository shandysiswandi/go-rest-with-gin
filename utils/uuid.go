package utils

import (
	"github.com/google/uuid"
)

// GenUUID is function to generate uuid for id table
func GenUUID() uuid.UUID {
	return uuid.New()
}
