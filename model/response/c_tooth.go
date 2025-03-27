package response

import "time"

type TeethRecorderHistoryResponse struct {
	CUserId        int             `json:"cUserId"`
	TeethRecorders []TeethRecorder `json:"teethRecorders"`
}

type TeethRecorder struct {
	ID       int       `json:"ID"`
	ExamDate time.Time `json:"examDate"`
	Examiner string    `json:"examiner"`
}
