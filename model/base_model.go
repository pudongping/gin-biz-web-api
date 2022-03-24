// package model 模型通用属性和方法
package model

import (
	"github.com/spf13/cast"
)

// BaseModel 模型基类
type BaseModel struct {
	// 主键 id  bigint(20) unsigned is_nullable NO
	ID uint `gorm:"primaryKey;autoIncrement;" json:"id"`
}

// CommonTimestampsField 时间戳字段
type CommonTimestampsField struct {
	// 创建时间  int(11) unsigned is_nullable NO
	CreatedAt int64 `json:"created_at"`
	// 更新时间  int(11) unsigned is_nullable NO
	UpdatedAt int64 `json:"updated_at"`
}

// GetStringID 获取 ID 的字符串格式
func (b BaseModel) GetStringID() string {
	return cast.ToString(b.ID)
}
