package service

import "mikrotik-script-generator/internal/model"

type HomeService interface {
	GetMikrotikScripts() []model.MikrotikScript
}

type homeService struct{}

func NewHomeService() HomeService {
	return &homeService{}
}

func (_self *homeService) GetMikrotikScripts() []model.MikrotikScript {
	return []model.MikrotikScript{
		{
			Name:        "Configure PPPoE",
			Description: "Set up PPPoE connection",
			Route:       "/pppoe",
		},
		{
			Name:        "Configure WireGuard",
			Description: "Set up WireGuard VPN configuration",
			Route:       "/wireguard",
		},
		{
			Name:        "Configure Multiple PPPoE Sessions with ECMP",
			Description: "Set up multiple PPPoE sessions on the same account and configure Equal-Cost Multi-Path (ECMP) routing for load balancing across those sessions",
			Route:       "/ecmp",
		},
		{
			Name:        "Configure IP Routing",
			Description: "Route specific IP addresses through a custom gateway",
			Route:       "/ip-routing",
		},
	}
}
