package service

import (
	"bytes"
	"html/template"
)

type PPPoEScriptService interface {
	GenerateScript(
		username,
		password string,
		interfaceName,
		lanNetwork string,
	) (string, error)
}

type pppoeScriptService struct{}

func NewPPPoEScriptService() PPPoEScriptService {
	return &pppoeScriptService{}
}

func (_self *pppoeScriptService) GenerateScript(
	username,
	password string,
	interfaceName,
	lanNetwork string,
) (string, error) {
	tmpl, err := template.ParseFiles("internal/service/mikrotik/pppoe_script.tmpl")
	if err != nil {
		return "", err
	}

	var script bytes.Buffer

	err = tmpl.Execute(&script, map[string]interface{}{
		"Username":   username,
		"Password":   password,
		"Interface":  interfaceName,
		"LANNetwork": lanNetwork,
	})
	if err != nil {
		return "", err
	}

	return script.String(), nil
}
