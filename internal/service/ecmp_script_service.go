package service

import (
	"bytes"
	"html/template"
)

type ECMPScriptService interface {
	GenerateScript(
		username,
		password string,
		sessions []int,
		interfaceName,
		lanNetwork string,
	) (string, error)
}

type ecmpScriptService struct{}

func NewECMPScriptService() ECMPScriptService {
	return &ecmpScriptService{}
}

func (_self *ecmpScriptService) GenerateScript(
	username,
	password string,
	sessions []int,
	interfaceName,
	lanNetwork string,
) (string, error) {
	tmpl, err := template.ParseFiles("internal/service/mikrotik/ecmp_script.tmpl")
	if err != nil {
		return "", err
	}

	var script bytes.Buffer

	err = tmpl.Execute(&script, map[string]interface{}{
		"Username":   username,
		"Password":   password,
		"Sessions":   sessions,
		"Interface":  interfaceName,
		"LANNetwork": lanNetwork,
	})
	if err != nil {
		return "", err
	}

	return script.String(), nil
}
