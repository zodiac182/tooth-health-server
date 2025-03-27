package v1

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zodiac182/tooth-health/server/core/logger"
	"github.com/zodiac182/tooth-health/server/global"
	"github.com/zodiac182/tooth-health/server/internal"
	"github.com/zodiac182/tooth-health/server/model/response"
	"github.com/zodiac182/tooth-health/server/model/system"
	"github.com/zodiac182/tooth-health/server/service"
)

type SysUserApi struct{}

// 获取用户列表
func (api *SysUserApi) ListSysUsers(c *gin.Context) {
	logger.Debug("ListSysUsers API called")

	username := strings.TrimSpace(c.Query("username"))
	nickname := strings.TrimSpace(c.Query("nickname"))
	if username != "" || nickname != "" {
		// 当前请求是通过用户名或姓名进行查找

		var responseData response.SysUsersResponse

		if username != "" {
			var userList []system.SysUser
			// 通过用户名进行查找，有且只有唯一的结果
			var user *system.SysUser
			user, err := service.UserServiceApp.GetUserByUserName(username)
			if err != nil {
				response.FailWithMessage(err.Error(), c)
				return
			}
			if user == nil {
				response.FailWithMessage("用户不存在", c)
				return
			}
			userList = append(userList, *user)
			responseData.Data = userList
			responseData.Total = 1

		} else {
			// 通过姓名进行查找，可能有多个结果
			userList, err := service.UserServiceApp.GetUserByNickname(nickname)
			if err != nil {
				response.FailWithMessage(err.Error(), c)
				return
			}
			responseData.Data = *userList
			responseData.Total = int(len(*userList))
		}

		response.OkWithData(responseData, c)
		return
	}

	pageStr := strings.TrimSpace(c.Query("page"))
	sizeStr := strings.TrimSpace(c.Query("size"))
	logger.Debug("pageStr: %s, sizeStr: %s", pageStr, sizeStr)

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		logger.Error("page is not a number. %+v", err)
		response.FailWithMessage(err.Error(), c)
		return
	}

	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		logger.Error("size is not a number. %+v", err)
		response.FailWithMessage(err.Error(), c)
		return
	}

	var userList *[]system.SysUser

	userList, total, err := service.UserServiceApp.GetAllUsers(page, size)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	var responseData response.SysUsersResponse
	responseData.Data = *userList
	responseData.Total = int(total)

	response.OkWithData(responseData, c)
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
