package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/zodiac182/tooth-health/server/core/logger"
	"github.com/zodiac182/tooth-health/server/model/request"
	"github.com/zodiac182/tooth-health/server/model/response"
	"github.com/zodiac182/tooth-health/server/model/system"
	"github.com/zodiac182/tooth-health/server/service"
)

type ToothApi struct{}

// 获取历史检查报告
func (t *ToothApi) GetRecordHistory(ctx *gin.Context) {
	logger.Debug("GetRecordHistory API called")

	var req *request.TeethRecorderHistoryRequest

	userId := req.ID

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", ctx)
		return
	}

	if req.ID == 0 {
		if req.IdCard == "" {
			response.FailWithMessage("参数错误", ctx)
			return
		} else {
			var cUser *system.CUser
			cUser, err := service.CUserServiceApp.GetUserByIdCard(req.IdCard)
			if err != nil {
				response.FailWithMessage("用户不存在", ctx)
				return
			}
			userId = int(cUser.ID)
		}
	}

	records, err := service.ToothServiceApp.GetToothRecordHistory(userId)
	if err != nil {
		response.FailWithMessage("获取历史记录失败", ctx)
		return
	}

	response.OkWithData(records, ctx)
}

// 创建报告，如果force创建，那么会覆盖之前的报告
func (t *ToothApi) CreateTeethRecord(c *gin.Context) {
	logger.Debug("CreateTeethRecord API called")
	var reqData request.TeethRecorderRequest
	if err := c.ShouldBindJSON(&reqData); err != nil {
		response.FailWithMessage("信息填写有误", c)
	}
	// 处理牙齿检查记录
	force := reqData.Force // 是否强制创建，即是否覆盖之前的记录

	// 将json里面的数据转换为ToothStatus切片，
	// 最终以json形式存数据库
	toothData := make([]system.ToothStatus, len(reqData.TeethData))
	for i, status := range reqData.TeethData {
		toothData[i] = system.ToothStatus{
			ToothID: status.ToothId,
			Status:  status.Status,
		}
	}

	teethRecord := system.TeethRecord{
		CUserID:   uint(reqData.UserId),
		TeethData: toothData,
		TeethExtraData: func(intSlice []int) pq.Int64Array {
			// 创建一个新的 int64 切片
			int64Slice := make([]int64, len(intSlice))

			// 将每个 int 转换为 int64
			for i, v := range intSlice {
				int64Slice[i] = int64(v)
			}

			// 转换为 pq.Int64Array
			return pq.Int64Array(int64Slice)
		}(reqData.TeethExtraData),
		Examiner: reqData.Examiner,
	}

	existed, err := service.CUserServiceApp.CreateOrUpdateTeethRecord(&teethRecord, force)
	if err != nil {
		response.FailWithMessage("创建牙齿检查记录失败", c)
		return
	}
	if existed {
		// 已经存在，返回code是7
		response.Existed(c)
		return
	}

	response.OkWithMessage("牙齿检查记录创建成功", c)

}
