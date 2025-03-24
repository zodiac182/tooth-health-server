package request

type SysUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	NickName string `json:"nickName"`
}
