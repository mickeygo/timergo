package entities

import (
	"github.com/satori/go.uuid"
)

// NewID 生成新的 uuid
func NewID() string {
	return uuid.NewV4().String()
}