package system

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

// CUser 结构体
// 包含用户信息、牙齿状态、复查信息
type CUser struct {
	gorm.Model
	IdCard        string      `json:"idCard" gorm:"unique;not null"`
	Name          string      `json:"name"`
	Gender        int         `json:"gender"` // 0: male, 1: female
	Phone         string      `json:"phone" binding:"min=8,max=11"`
	School        string      `json:"school"`
	Class         string      `json:"class"`
	TTeethRecords TeethRecord `gorm:"foreignkey:CUserID"`
}

func (CUser) TableName() string {
	return "c_users"
}

// TeethRecord 结构体
// 用于记录单次牙齿检查记录
type TeethRecord struct {
	gorm.Model
	CUserID uint `gorm:"not null;index"` // 外键，指向 CUser

	// 牙齿状态, 以json存放
	TeethData ToothStatusArray `gorm:"type:json"` // 例如 {"11": 1, "12": 0, "13": 2}

	TeethExtraData pq.Int64Array `gorm:"type:integer[]"` // 其他状态（需要引入 lib/pq）

	Examiner string `gorm:"size:20"`  //
	Comments string `gorm:"size:255"` // 备注信息
}

// ToothStatus 结构体
// 用于记录单个牙齿的状态
type ToothStatus struct {
	ToothID int `json:"id"`
	Status  int `json:"status"`
}

// 定义一个结构，用于将切片转换成 JSON 存储
type ToothStatusArray []ToothStatus

// `Value` 方法：将 ToothStatusArray 转换成 JSON 存储
func (t ToothStatusArray) Value() (driver.Value, error) {
	return json.Marshal(t)
}

// `Scan` 方法：从数据库 JSON 解析回 ToothStatusArray
func (t *ToothStatusArray) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), t)
}

func (TeethRecord) TableName() string {
	return "teeth_records"
}
