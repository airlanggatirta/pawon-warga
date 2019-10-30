package common

import (
	"bytes"
	"text/template"
)

func LoadTemplate(filePath string, data interface{}) (string, error) {
	template, err := template.ParseFiles(filePath)
	if err != nil {
		return "", err
	}

	var buff bytes.Buffer
	err = template.Execute(&buff, data)
	if err != nil {
		return "", nil
	}

	return buff.String(), nil
}
