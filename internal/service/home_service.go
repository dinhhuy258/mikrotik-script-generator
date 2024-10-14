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
		{Name: "Configure WireGuard", Description: "Set up WireGuard VPN configuration", Route: "/wireguard"},
	}
}
