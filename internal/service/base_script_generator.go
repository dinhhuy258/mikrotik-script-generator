package service

import (
	"bytes"
	"html"
	"html/template"
)

type BaseScriptGenerator struct{}

func (_self BaseScriptGenerator) GenerateScriptFromTemplate(
	templateFile string,
	data map[string]any,
) (string, error) {
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return "", err
	}

	var script bytes.Buffer

	err = tmpl.Execute(&script, data)
	if err != nil {
		return "", err
	}

	return html.UnescapeString(script.String()), nil
}
