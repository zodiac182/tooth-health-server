package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zodiac182/tooth-health/server/middleware"
	"github.com/zodiac182/tooth-health/server/model/request"
	"github.com/zodiac182/tooth-health/server/model/response"
	"github.com/zodiac182/tooth-health/server/service"
)

type LoginApi struct{}

func (l *LoginApi) Login(c *gin.Context) {
	var loginReq request.SysLoginReq
	var loginResp response.SysLoginResp
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		response.FailWithMessage("用户名和密码不能为空", c)
		return
	}

	user, error := service.LoginServiceApp.Login(&loginReq)
	if error != nil {
		// 登录失败，用户名或者密码错误
		response.FailWithMessage(error.Error(), c)
		return
		// handle error
	}
	loginResp.NickName = user.Nickname
	loginResp.UserName = user.Username
	token, err := middleware.GenerateJwt(user)
	if err != nil {
		// handle error
		response.FailWithMessage("生成token失败", c)
		return
	}
	loginResp.Token = token
	loginResp.Role = user.Role

	response.OkWithDetails(loginResp, "登陆成功", c)
}
