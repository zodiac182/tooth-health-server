package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zodiac182/tooth-health/server/api"
	"github.com/zodiac182/tooth-health/server/global"
	"github.com/zodiac182/tooth-health/server/middleware"
)

//

func Routers() *gin.Engine {
	r := gin.Default()

	var routerGroupApp = new(api.ApiGroup)
	r.Use(middleware.CORSMiddleware())
	r.GET("/version", routerGroupApp.VersionApi.GetVersion)

	// 注册public路由
	publicGroup := r.Group(global.RouterPrefix)

	{
		publicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
		publicGroup.POST("/login", routerGroupApp.LoginApi.Login)
	}

	// 注册protected路由
	protectedGroup := r.Group(global.RouterPrefix)
	protectedGroup.Use(middleware.JwtMiddleware())
	{
		// 系统用户相关
		protectedGroup.GET("/version", routerGroupApp.VersionApi.GetVersion)
		protectedGroup.GET("/sysusers", routerGroupApp.SysUserApi.ListSysUsers)                    // 用户列表
		protectedGroup.POST("/sysuser", routerGroupApp.SysUserApi.AddSysUser)                      // 创建用户
		protectedGroup.PUT("/sysuser/role", routerGroupApp.SysUserApi.UpdateSysUserRole)           // 修改用户角色
		protectedGroup.PUT("/sysuser", routerGroupApp.SysUserApi.UpdateSysUserInfo)                // 修改密码
		protectedGroup.DELETE("/sysuser/:id", routerGroupApp.SysUserApi.DeleteSysUser)             // 删除用户
		protectedGroup.PUT("/sysuser/reset/password/:id", routerGroupApp.SysUserApi.ResetPassword) // 修改密码

	}

	{
		// 用户相关
		protectedGroup.POST("/cuser", routerGroupApp.CUserApi.CreateOrUpdateCUser)
		protectedGroup.GET("/cuser", routerGroupApp.CUserApi.GetUserByIdCard) // 通过身份证号获取用户信息
	}

	{
		// 检查记录相关
		protectedGroup.POST("/tooth/record", routerGroupApp.ToothApi.CreateTeethRecord)   // 创建牙齿检查记录
		protectedGroup.GET("/tooth/record/all", routerGroupApp.ToothApi.GetRecordHistory) // 获取用户所有检查记录概要
	}
	return r
}
