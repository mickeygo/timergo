package repository

// 仓储接口定义
type Repository interface {
	// 查找实体对象
	Find(id string) interface{}

	// 添加实体对象
	Add(entity interface{})

	// 批量添加实体对象
	AddRange(entities interface{})

	// 更新实体对象
	Update(entity interface{}) bool

	// 批量更新实体对象
	UpdateRange(entities interface{})

	// 删除实体对象
	Delete(id string) bool

	// 批量删除实体对象
	DeleteRange(entities interface{}) bool
}