package service

import (
	"bytes"
	"html/template"
)

type PPPoEScriptService interface {
	GenerateScript(
		username,
		password,
		interfaceName string,
		bridgeLANPorts []string,
		gateway,
		lanNetwork,
		dhcpRange string,
	) (string, error)
}

type pppoeScriptService struct{}

func NewPPPoEScriptService() PPPoEScriptService {
	return &pppoeScriptService{}
}

func (_self *pppoeScriptService) GenerateScript(
	username,
	password,
	interfaceName string,
	bridgeLANPorts []string,
	gateway,
	lanNetwork,
	dhcpRange string,
) (string, error) {
	tmpl, err := template.ParseFiles("internal/service/mikrotik/pppoe_script.tmpl")
	if err != nil {
		return "", err
	}

	var script bytes.Buffer

	err = tmpl.Execute(&script, map[string]interface{}{
		"Username":       username,
		"Password":       password,
		"Interface":      interfaceName,
		"BridgeLANPorts": bridgeLANPorts,
		"Gateway":        gateway,
		"LANNetwork":     lanNetwork,
		"DHCPRange":      dhcpRange,
	})
	if err != nil {
		return "", err
	}

	return script.String(), nil
}
