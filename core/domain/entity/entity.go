package entity

import (
	"github.com/satori/go.uuid"
)

// 生成新的 Id
func NewId() string {
	return uuid.NewV4().String()
}