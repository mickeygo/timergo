package entity

import (
	"github.com/satori/go.uuid"
)

type Entity struct {
	// 实体 Id
	Id string `json:"id"`
}

func (e *Entity) NewId() {
	e.Id = uuid.NewV4().String()
}