package util

import "github.com/google/uuid"

func NewUuid() string {
	return uuid.NewString()
}

func IsUuid(id string) bool {
	_, err := uuid.Parse(id)

	return err == nil
}
