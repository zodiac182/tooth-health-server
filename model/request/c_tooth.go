package request

// 用户提交信息的请求体
type CUserRequest struct {
	IdCard string `json:"idCard"`
	Name   string `json:"name"`
	Gender int    `json:"gender"` // 0: male, 1: female
	Phone  string `json:"phone" binding:"min=8,max=11"`
	School string `json:"school"`
	Class  string `json:"class"`
}

type TeethRecorderRequest struct {
	UserId    int           `json:"userId"`
	RecordId  int           `json:"recordId"`
	TeethData []ToothStatus `json:"teethData"`
	//  { text: "牙齿不齐", value: 1 },
	// { text: "开合", value: 2 },
	// { text: "反合", value: 3 },
	// { text: "唇舌系带过短", value: 4 },
	// { text: "乳牙滞留", value: 5 },
	// { text: "口腔黏膜异常", value: 6 }
	TeethExtraData []int  `json:"teethExtraData"`
	Examiner       string `json:"examiner"`
	Force          bool   `json:"force"` // 是否强制提交
}

type ToothStatus struct {
	// { label: "无龋", value: 0, color: "white" },
	// { label: "有龋", value: 1, color: "red" },
	// { label: "已填充有龋", value: 2, color: "orange" },
	// { label: "已填充无龋", value: 3, color: "green" },
	// { label: "因龋缺失", value: 4, color: "gray" },
	ToothId int `json:"id"`     // 牙齿编号
	Status  int `json:"status"` // 检查状态 (0: 无龋, 1: 有龋, 2: 已填充有龋, etc.)
}

// 用于获取历史记录
type TeethRecorderHistoryRequest struct {
	ID     int    `json:"ID"`
	IdCard string `json:"idCard"`
}
