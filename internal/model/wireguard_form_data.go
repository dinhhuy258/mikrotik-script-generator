package model

import "mime/multipart"

type WireGuardFormData struct {
	Name       string                `form:"name"`
	ListenPort int                   `form:"listenPort"`
	ConfigType string                `form:"configType"`
	ConfigFile *multipart.FileHeader `form:"configFile"`
	Errors     map[string]string     `form:"errors"`
}
