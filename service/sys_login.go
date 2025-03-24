package service

import (
	"errors"

	"github.com/zodiac182/tooth-health/server/core/logger"
	"github.com/zodiac182/tooth-health/server/internal"
	"github.com/zodiac182/tooth-health/server/model/request"
	"github.com/zodiac182/tooth-health/server/model/system"
	"gorm.io/gorm"
)

type LoginService struct {
	db *gorm.DB
}

func NewLoginService(db *gorm.DB) *LoginService {
	return &LoginService{db: db}
}
func (s *LoginService) Login(loginInfo *request.SysLoginReq) (*system.SysUser, error) {
	logger.Debug("Login service called")
	var user *system.SysUser
	username := loginInfo.Username
	password := loginInfo.Password

	user, err := UserServiceApp.GetUserByUsername(username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}
	if !internal.CheckPassword(password, user.Password) {
		return nil, errors.New("密码错误")
	}
	return user, nil
}
