package system

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type CUser struct {
	gorm.Model
	IdCard       string      `json:"idCard" gorm:"unique;not null"`
	Name         string      `json:"name"`
	Gender       int         `json:"gender"` // 0: male, 1: female
	Phone        string      `json:"phone" binding:"min=8,max=11"`
	School       string      `json:"school"`
	Class        string      `json:"class"`
	TeethReports TeethReport `gorm:"foreignkey:CUserID"`
}

type TeethReport struct {
	gorm.Model
	CUserID uint `gorm:"not null;index"` // 外键，指向 CUser

	// 牙齿状态存成 JSON 或者拆分到 ToothCondition
	ToothStatuses ToothStatusArray `gorm:"type:json"` // 例如 {"11": 1, "12": 0, "13": 2}

	OtherStatus pq.Int64Array `gorm:"type:integer[]"` // 其他状态（需要引入 lib/pq）

	Checker  string `gorm:"size:20"`  // 复查人
	Comments string `gorm:"size:255"` // 备注信息
}
type ToothStatus struct {
	ToothID int `json:"id"`
	Status  int `json:"status"`
}

type ToothStatusArray []ToothStatus

// `Value` 方法：将 ToothStatusArray 转换成 JSON 存储
func (t ToothStatusArray) Value() (driver.Value, error) {
	return json.Marshal(t)
}

// `Scan` 方法：从数据库 JSON 解析回 ToothStatusArray
func (t *ToothStatusArray) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), t)
}
