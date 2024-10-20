package service

import (
	"bytes"
	"html/template"
)

type ECMPScriptService interface {
	GenerateScript(
		username,
		password string,
		numSessions int,
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
	numSessions int,
	interfaceName,
	lanNetwork string,
) (string, error) {
	tmpl, err := template.ParseFiles("internal/service/mikrotik/ecmp_script.tmpl")
	if err != nil {
		return "", err
	}

	var script bytes.Buffer

	sessions := make([]int, numSessions)
	for i := 0; i < numSessions; i++ {
		sessions[i] = i + 1
	}

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
