package service

import (
	"github.com/zodiac182/tooth-health/server/core/logger"
	"github.com/zodiac182/tooth-health/server/global"
	"github.com/zodiac182/tooth-health/server/internal"
	"github.com/zodiac182/tooth-health/server/model/system"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) CreateUser(user *system.SysUser) error {
	return s.db.Create(user).Error
}

// 通过用户名查找用户
func (s *UserService) GetUserByUserName(username string) (*system.SysUser, error) {
	logger.Debug("GetUserByUserName called: %s", username)
	var user system.SysUser
	err := s.db.First(&user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// 通过姓名查找用户
func (s *UserService) GetUserByNickname(nickname string) (*[]system.SysUser, error) {
	logger.Debug("GetUserByNickname called: %s", nickname)
	var users []system.SysUser
	err := s.db.Find(&users, "nickname LIKE ?", "%"+nickname+"%").Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (s *UserService) GetAllUsers(page, size int) (*[]system.SysUser, int64, error) {
	logger.Debug("GetAllUsers service called")
	// 先获取总数
	var total int64
	if err := s.db.Model(&system.SysUser{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var users []system.SysUser

	// 计算offset
	offset := (page - 1) * size

	err := s.db.Offset(offset).Limit(size).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}
	return &users, total, nil
}

func (s *UserService) AddUser(user *system.SysUser) error {
	logger.Debug("AddUser called: %s", user.Username)
	return s.db.Create(user).Error
}

func (s *UserService) UpdateUserRole(user *system.SysUser) error {
	logger.Debug("UpdateUserRole service called: %s", user.Username)

	return s.db.Model(&system.SysUser{}).Where("username = ?", user.Username).Update("role", user.Role).Error
}

func (s *UserService) DeleteUser(user *system.SysUser) error {
	logger.Debug("DeleteUser service called: %s", user.Username)
	return s.db.Unscoped().Delete(user).Error
}

// 重置密码为123456
func (s *UserService) ResetPassword(user *system.SysUser) error {
	logger.Debug("ResetPassword service called: %s", user.ID)
	pwd, err := internal.HashPassword(global.OriginalPassword)
	if err != nil {
		return err
	}
	return s.db.Model(&system.SysUser{}).Where("id = ?", user.ID).Update("password", pwd).Error
}

// 修改用户信息，包括用户姓名和密码
func (s *UserService) UpdateUserInfo(user *system.SysUser) error {
	logger.Debug("UpdateUserInfo service called: %s", user.Username)
	logger.Debug("UpdateUserInfo service called: %s", user.Password)
	logger.Debug("UpdateUserInfo service called: %s", user.Nickname)
	if user.Password != "" {
		pwd, err := internal.HashPassword(user.Password)
		if err != nil {
			return err
		}
		user.Password = pwd
	}
	return s.db.Model(&system.SysUser{}).Where("username= ?", user.Username).Updates(user).Error
}
