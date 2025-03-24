package response

import "gorm.io/gorm"

type NewCUserResponse struct {
	gorm.Model
	Code    int    `json:"code"`
	Message string `json:"message"`
}
