package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zodiac182/tooth-health/server/global"
	"github.com/zodiac182/tooth-health/server/model/response"
)

// SysVersion 结构体用于处理与系统版本相关的请求
type VersionApi struct{}

// @Summary 获取系统版本
// @Description 返回当前系统的版本号
// @Tags 系统信息
// @Accept json
// @Produce json
// @Success 200 {object} gin.H "成功返回版本号"
// @Router /version [get]
func (v *VersionApi) GetVersion(c *gin.Context) {
	response.OkWithMessage("当前系统版本号为"+global.Version, c)
}
