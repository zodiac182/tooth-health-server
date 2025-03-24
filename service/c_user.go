package service

import (
	"github.com/zodiac182/tooth-health/server/core/logger"
	"github.com/zodiac182/tooth-health/server/model/system"
	"gorm.io/gorm"
)

type CUserService struct {
	db *gorm.DB
}

func NewCUserService(db *gorm.DB) *CUserService {
	return &CUserService{db: db}
}

func (u *CUserService) CreateOrUpdateUser(user *system.CUser) error {
	existingCUser, err := u.GetUserByIdCard(user.IdCard)
	if err == nil {
		// 已存在用户，更新用户信息
		if existingCUser.Name != user.Name || existingCUser.Phone != user.Phone || existingCUser.School != user.School || existingCUser.Class != user.Class {
			return u.UpdateUser(user)
		} else {
			user.ID = existingCUser.ID
			return nil
		}
	}
	// 新增用户
	return u.db.Create(user).Error
}

func (u *CUserService) GetUserByName(name string) error {
	logger.Debug("GetUserByName called: %s", name)
	var user system.CUser
	err := u.db.Where("name = ?", name).First(&user).Error
	if err != nil {
		return err
	}
	return nil
}

// @todo: add GetUserByPhone
func (u *CUserService) GetUserByPhone(phone string) error {
	logger.Debug("GetUserByPhone called: %s", phone)
	var user system.CUser
	err := u.db.Where("phone = ?", phone).First(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *CUserService) GetUserByIdCard(idCard string) (*system.CUser, error) {
	logger.Debug("GetUserB	yPhone called: %s", idCard)
	var user system.CUser
	err := u.db.Where("id_card = ?", idCard).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *CUserService) UpdateUser(user *system.CUser) error {
	return u.db.Save(user).Error
}

func (u *CUserService) CreateTeethReport(report *system.TeethReport) error {
	return u.db.Create(report).Error
}
