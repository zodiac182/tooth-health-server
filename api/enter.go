package api

import v1 "github.com/zodiac182/tooth-health/server/api/v1"

type ApiGroup struct {
	VersionApi v1.VersionApi
	LoginApi   v1.LoginApi
	SysUserApi v1.SysUserApi
	CUserApi   v1.CUserApi
}
