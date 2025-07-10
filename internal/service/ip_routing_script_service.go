package service

import (
	"mikrotik-script-generator/internal/model"
)

type IPRoutingScriptService interface {
	GenerateScript(formData model.IPRoutingFormData) (string, error)
	GenerateReverseScript(formData model.IPRoutingFormData) (string, error)
}

type ipRoutingScriptService struct {
	BaseScriptGenerator
}

func NewIPRoutingScriptService() IPRoutingScriptService {
	return &ipRoutingScriptService{}
}

func (_self *ipRoutingScriptService) GenerateScript(formData model.IPRoutingFormData) (string, error) {
	data := map[string]any{
		"IPAddresses":  formData.IPAddresses,
		"Gateway":      formData.Gateway,
		"RoutingTable": formData.RoutingTable,
	}

	return _self.GenerateScriptFromTemplate("internal/service/mikrotik/ip_routing_script.tmpl", data)
}

func (_self *ipRoutingScriptService) GenerateReverseScript(formData model.IPRoutingFormData) (string, error) {
	data := map[string]any{
		"IPAddresses":  formData.IPAddresses,
		"Gateway":      formData.Gateway,
		"RoutingTable": formData.RoutingTable,
	}

	return _self.GenerateScriptFromTemplate("internal/service/mikrotik/ip_routing_reverse_script.tmpl", data)
}
