package logic

import (
	"github.com/google/uuid"
)

func NewUUID() UUID {
	return UUID(uuid.New().String())
}
