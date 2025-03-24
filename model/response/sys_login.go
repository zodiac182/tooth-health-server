package response

type SysLoginResp struct {
	Token    string `json:"token"`
	NickName string `json:"nickname"`
	UserName string `json:"username"`
	Role     string `json:"role"`
}
