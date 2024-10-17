package model

type PPPoEFormData struct {
	Username   string `form:"username"`
	Password   string `form:"password"`
	Interface  string `form:"interface"`
	LANNetwork string `form:"lanNetwork"`
}
