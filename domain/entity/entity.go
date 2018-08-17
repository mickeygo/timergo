package entity

import (
	"github.com/satori/go.uuid"
)

// 实体对象
type Entity struct {
	// 实体 Id
	Id string `json:"id"`
}

// 生成新的 uuid
func (e *Entity) NewId() {
	e.Id = uuid.NewV4().String()
}