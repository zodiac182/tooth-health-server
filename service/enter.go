package service

import (
	"github.com/zodiac182/tooth-health/server/core/logger"
	"github.com/zodiac182/tooth-health/server/db"
)

var UserServiceApp *UserService
var LoginServiceApp *LoginService
var CUserServiceApp *CUserService

func InitService() {
	UserServiceApp = NewUserService(db.DB)
	LoginServiceApp = NewLoginService(db.DB)
	CUserServiceApp = NewCUserService(db.DB)
	logger.Debug("All services initialized")
}
