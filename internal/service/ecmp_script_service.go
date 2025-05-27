package service

type ECMPScriptService interface {
	GenerateScript(
		username,
		password string,
		numSessions int,
		interfaceName,
		lanNetwork string,
	) (string, error)
}

type ecmpScriptService struct {
	BaseScriptGenerator
}

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
	sessions := make([]int, numSessions)
	for i := 0; i < numSessions; i++ {
		sessions[i] = i + 1
	}

	return _self.GenerateScriptFromTemplate("internal/service/mikrotik/ecmp_script.tmpl", map[string]any{
		"Username":   username,
		"Password":   password,
		"Sessions":   sessions,
		"Interface":  interfaceName,
		"LANNetwork": lanNetwork,
	})
}
