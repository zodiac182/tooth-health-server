package service

import (
	"time"

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
	logger.Debug("CreateOrUpdateUser service called: %v", user)
	existingCUser, err := u.GetUserByIdCard(user.IdCard)
	if err == nil {
		logger.Debug("用户已存在")
		user.ID = existingCUser.ID
		// 已存在用户，更新用户信息
		if existingCUser.Name != user.Name || existingCUser.Phone != user.Phone ||
			existingCUser.School != user.School || existingCUser.Class != user.Class ||
			existingCUser.Gender != user.Gender {
			logger.Debug("需要更新用户信息")

		} else {
			return nil
		}
	}
	// 新增用户
	return u.UpdateUser(user)
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

// 通过身份证号获取信息
func (u *CUserService) GetUserByIdCard(idCard string) (*system.CUser, error) {
	logger.Debug("GetUserByIdCard called: %s", idCard)
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

func (u *CUserService) CreateOrUpdateTeethRecord(record *system.TeethRecord, force bool) (existed bool, err error) {
	// 检查当日是否已经有记录
	today := time.Now().Format("2006-01-02")
	// start := today + "00:00:00"
	// end := today + "23:59:59"
	existed = false
	var currRecord *system.TeethRecord
	err = u.db.Where("DATE(created_at) = ?", today).Where("c_user_id = ?", record.CUserID).First(&currRecord).Error
	if err == nil {
		// 记录已经存在
		if !force {
			// 不强制更新 直接返回
			return true, nil
		} else {
			// 强制更新
			record.ID = currRecord.ID
			return false, u.db.Save(record).Error
		}
	}

	return false, u.db.Create(record).Error
}
