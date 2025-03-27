package service

import (
	"github.com/zodiac182/tooth-health/server/core/logger"
	"github.com/zodiac182/tooth-health/server/model/response"
	"github.com/zodiac182/tooth-health/server/model/system"
	"gorm.io/gorm"
)

type ToothService struct {
	db *gorm.DB
}

func NewToothService(db *gorm.DB) *ToothService {
	return &ToothService{db: db}
}

// 获取牙齿检查记录历史
func (t *ToothService) GetToothRecordHistory(id int) (*response.TeethRecorderHistoryResponse, error) {
	logger.Debug("GetToothRecordHistory service called. id: %d", id)

	var data *[]system.TeethRecord
	if err := t.db.Where("c_user_id =?", id).Find(&data).Error; err != nil {
		logger.Error("GetToothRecordHistory failed. err: %v", err)
		return nil, err
	}

	// 组装返回类型
	var teethRecorders []response.TeethRecorder
	for _, item := range *data {
		teethRecorders = append(teethRecorders, response.TeethRecorder{
			ID:       int(item.ID),
			ExamDate: item.CreatedAt,
			Examiner: item.Examiner,
		})
	}

	var result = &response.TeethRecorderHistoryResponse{
		CUserId:        id,
		TeethRecorders: teethRecorders,
	}
	return result, nil
}
