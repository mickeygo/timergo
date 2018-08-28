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

// 检查实体是否是临时的
func (e *Entity) IsTransient() bool {
	return e.Id == ""
}

// 比较两个实体对象是否相同
func Equal(a interface{}, b interface{}) bool {
	e1, ok := a.(Entity)
	if (!ok) {
		return false
	}

	e2, ok := b.(Entity)
	if (!ok) {
		return false
	}

	return e1.Id == e2.Id
}