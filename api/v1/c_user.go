package v1

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zodiac182/tooth-health/server/core/logger"
	"github.com/zodiac182/tooth-health/server/model/request"
	"github.com/zodiac182/tooth-health/server/model/response"
	"github.com/zodiac182/tooth-health/server/model/system"
	"github.com/zodiac182/tooth-health/server/service"
)

type CUserApi struct{}

// 创建用户信息，如果用户已经存在，则更新用户信息
func (api *CUserApi) CreateOrUpdateCUser(c *gin.Context) {
	logger.Debug("CreateOrUpdateCUser API called")
	var reqData request.CUserRequest
	if err := c.ShouldBindJSON(&reqData); err != nil {
		response.FailWithMessage("信息填写有误", c)
	}

	var cUser system.CUser

	cUser.IdCard = reqData.IdCard
	cUser.Name = reqData.Name
	cUser.Gender = reqData.Gender
	cUser.Phone = reqData.Phone
	cUser.School = reqData.School
	cUser.Class = reqData.Class

	// 处理用户
	err := service.CUserServiceApp.CreateOrUpdateUser(&cUser)
	if err != nil {
		response.FailWithMessage("用户信息保存失败", c)
	}

	response.OkWithDetails(cUser, "用户信息保存成功", c)
}

// 通过身份证号获取用户信息
func (api *CUserApi) GetUserByIdCard(c *gin.Context) {
	logger.Debug("GetUserByIdCard API called")
	idCard := strings.TrimSpace(c.Query("idCard"))
	if idCard == "" {
		response.FailWithMessage("身份证号不能为空", c)
		return
	}

	cUser, err := service.CUserServiceApp.GetUserByIdCard(idCard)
	if err != nil {
		response.FailWithMessage("获取用户信息失败", c)
		return
	}

	response.OkWithData(cUser, c)
}
