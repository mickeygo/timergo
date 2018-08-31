package entity

import (
	"github.com/satori/go.uuid"
)

// NewID 生成新的 Id
func NewID() string {
	return uuid.NewV4().String()
}