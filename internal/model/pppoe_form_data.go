package model

type PPPoEFormData struct {
	Username      string `form:"username"`
	Password      string `form:"password"`
	Interface     string `form:"interface"`
	BridgeLANPort string `form:"bridgeLANPort"`
	Gateway       string `form:"gateway"`
	LANNetwork    string `form:"lanNetwork"`
	DHCPRange     string `form:"dhcpRange"`
}
