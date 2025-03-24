package v1

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zodiac182/tooth-health/server/core/logger"
	"github.com/zodiac182/tooth-health/server/global"
	"github.com/zodiac182/tooth-health/server/internal"
	"github.com/zodiac182/tooth-health/server/model/response"
	"github.com/zodiac182/tooth-health/server/model/system"
	"github.com/zodiac182/tooth-health/server/service"
)

type SysUserApi struct{}

func (api *SysUserApi) ListSysUsers(c *gin.Context) {
	var userList *[]system.SysUser

	userList, err := service.UserServiceApp.GetAllUsers()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(userList, c)
}

// 添加用户
func (api *SysUserApi) AddSysUser(c *gin.Context) {
	userInfo := system.SysUser{}
	if err := c.ShouldBindJSON(&userInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 初始密码
	pwd, err := internal.HashPassword(global.OriginalPassword)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	userInfo.Password = pwd

	err = service.UserServiceApp.AddUser(&userInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("添加成功", c)
}

// 更改用户角色
func (api *SysUserApi) UpdateSysUserRole(c *gin.Context) {
	logger.Debug("UpdateSysUserRole API called")
	var sysUser system.SysUser
	if err := c.ShouldBindJSON(&sysUser); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := service.UserServiceApp.UpdateUserRole(&sysUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// 删除用户
func (api *SysUserApi) DeleteSysUser(c *gin.Context) {
	logger.Debug("DeleteSysUser API called")
	userIdStr := c.Param("id")
	var sysUser system.SysUser
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	sysUser.ID = uint(userId)
	err = service.UserServiceApp.DeleteUser(&sysUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// 重置密码
func (api *SysUserApi) ResetPassword(c *gin.Context) {
	logger.Debug("ResetPassword API called")
	userIdStr := c.Param("id")
	var sysUser system.SysUser
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	sysUser.ID = uint(userId)
	err = service.UserServiceApp.ResetPassword(&sysUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("重置密码成功", c)
}

// 修改用户信息
func (api *SysUserApi) UpdateSysUserInfo(c *gin.Context) {
	logger.Debug("UpdateSysUserInfo API called")
	var sysUser system.SysUser
	if err := c.ShouldBindJSON(&sysUser); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := service.UserServiceApp.UpdateUserInfo(&sysUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}
