// package model 模型通用属性和方法
package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"time"

	"github.com/spf13/cast"
)

type TimeNormal struct {
	time.Time
}

// MarshalJSON 将时间字段转换为 `%Y-%m-%d %H:%M:%S` 格式
func (t TimeNormal) MarshalJSON() ([]byte, error) {
	if y := t.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("TimeNormal.MarshalJSON: year outside of range [0,9999]")
	}
	tune := t.Format(`"2006-01-02 15:04:05"`)
	return []byte(tune), nil
}

// Value 写入 mysql 时，需要调用这个方法
func (t TimeNormal) Value() (driver.Value, error) {
	var zeroTime time.Time
	// 判断给定时间是否和默认零时间的时间戳相同
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan 在数据查询出来之前对数据进行操作
func (t *TimeNormal) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = TimeNormal{Time: value}
		return nil
	}

	return fmt.Errorf("can not convert %v to timestamp", value)
}

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
