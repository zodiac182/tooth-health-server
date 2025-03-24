package system

import "gorm.io/gorm"

const (
	AdminRole = "admin"
	UserRole  = "user"
)

type SysUser struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Nickname string `json:"nickname" gorm:"default: 用户"`
	Role     string `json:"role" gorm:"default: user"`
}

func (SysUser) TableName() string {
	return "sys_users" // 明确指定表名
}
