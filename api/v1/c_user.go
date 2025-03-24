package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/zodiac182/tooth-health/server/model/request"
	"github.com/zodiac182/tooth-health/server/model/response"
	"github.com/zodiac182/tooth-health/server/model/system"
	"github.com/zodiac182/tooth-health/server/service"
)

type CUserApi struct{}

func (api *CUserApi) CreateUserAndTeethCheck(c *gin.Context) {
	var reqData request.CTeethReport
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

	var existingCUser *system.CUser
	existingCUser, err := service.CUserServiceApp.GetUserByIdCard(reqData.IdCard)
	if err != nil {
		err = service.CUserServiceApp.CreateOrUpdateUser(&cUser)
		if err != nil {
			response.FailWithMessage("创建用户失败", c)
		}
	} else {
		cUser.ID = existingCUser.ID
	}

	toothStatuses := make([]system.ToothStatus, len(reqData.TeethStatus))
	for i, status := range reqData.TeethStatus {
		toothStatuses[i] = system.ToothStatus{
			ToothID: status.ToothId,
			Status:  status.Status,
		}
	}

	teethReport := system.TeethReport{
		CUserID:       cUser.ID,
		ToothStatuses: toothStatuses,
		OtherStatus: func(intSlice []int) pq.Int64Array {
			// 创建一个新的 int64 切片
			int64Slice := make([]int64, len(intSlice))

			// 将每个 int 转换为 int64
			for i, v := range intSlice {
				int64Slice[i] = int64(v)
			}

			// 转换为 pq.Int64Array
			return pq.Int64Array(int64Slice)
		}(reqData.TeethOtherStatus),
		Checker: reqData.Checker,
	}

	err = service.CUserServiceApp.CreateTeethReport(&teethReport)
	if err != nil {
		response.FailWithMessage("创建牙齿检查记录失败", c)
	}

	response.OkWithMessage("牙齿检查记录创建成功", c)

}
