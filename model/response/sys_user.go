package response

import "github.com/zodiac182/tooth-health/server/model/system"

type SysUsersResponse struct {
	Data  []system.SysUser `json:"data"`
	Total int              `json:"total"`
}
