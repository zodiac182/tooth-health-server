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
		protectedGroup.GET("/version", routerGroupApp.VersionApi.GetVersion)
		protectedGroup.GET("/sysusers", routerGroupApp.SysUserApi.ListSysUsers)                    // 用户列表
		protectedGroup.POST("/sysuser", routerGroupApp.SysUserApi.AddSysUser)                      // 创建用户
		protectedGroup.PUT("/sysuser/role", routerGroupApp.SysUserApi.UpdateSysUserRole)           // 修改用户角色
		protectedGroup.PUT("/sysuser", routerGroupApp.SysUserApi.UpdateSysUserInfo)                // 修改密码
		protectedGroup.DELETE("/sysuser/:id", routerGroupApp.SysUserApi.DeleteSysUser)             // 删除用户
		protectedGroup.PUT("/sysuser/reset/password/:id", routerGroupApp.SysUserApi.ResetPassword) // 修改密码

	}

	{
		protectedGroup.POST("/report/new", routerGroupApp.CUserApi.CreateUserAndTeethCheck)
	}

	return r
}
