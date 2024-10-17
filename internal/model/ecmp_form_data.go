package model

type ECMPFormData struct {
	Username   string `form:"username"`
	Password   string `form:"password"`
	Sessions   int    `form:"sessions"`
	Interface  string `form:"interface"`
	LANNetwork string `form:"lanNetwork"`
}
