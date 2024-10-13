package service

import "mikrotik-script-generator/internal/model"

type HomeService interface {
	GetScripts() []model.Script
}

type homeService struct{}

func NewHomeService() HomeService {
	return &homeService{}
}

func (_self *homeService) GetScripts() []model.Script {
	return []model.Script{
		{Name: "Configure WireGuard", Description: "Set up WireGuard VPN configuration", Route: "/wireguard"},
	}
}
