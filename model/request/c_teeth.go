package request

type CTeethReport struct {
	IdCard      string        `json:"idCard" gorm:"unique;not null"`
	Name        string        `json:"name"`
	Gender      int           `json:"gender"` // 0: male, 1: female
	Phone       string        `json:"phone" binding:"min=8,max=11"`
	School      string        `json:"school"`
	Class       string        `json:"class"`
	Checker     string        `json:"checker"`
	TeethStatus []ToothStatus `json:"teethStatus""`
	//  { text: "牙齿不齐", value: 1 },
	// { text: "开合", value: 2 },
	// { text: "反合", value: 3 },
	// { text: "唇舌系带过短", value: 4 },
	// { text: "乳牙滞留", value: 5 },
	// { text: "口腔黏膜异常", value: 6 }
	TeethOtherStatus []int `json:"teethOtherStatus"` //
}

type ToothStatus struct {
	// { label: "无龋", value: 0, color: "white" },
	// { label: "有龋", value: 1, color: "red" },
	// { label: "已填充有龋", value: 2, color: "orange" },
	// { label: "已填充无龋", value: 3, color: "green" },
	// { label: "因龋缺失", value: 4, color: "gray" },
	ToothId int `json:"id" gorm:"not null"`     // 牙齿编号
	Status  int `json:"status" gorm:"not null"` // 检查状态 (0: 无龋, 1: 有龋, 2: 已填充有龋, etc.)
}
