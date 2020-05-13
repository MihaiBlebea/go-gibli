package builder

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/MihaiBlebea/go-gibli/transformer"
)

// BuildModel receives the data, file to write to and template path and writes the model to a go file
func BuildModel(model Model, path string) error {
	functionMap := template.FuncMap{
		"toUpperCase":    strings.Title,
		"toLowerCase":    strings.ToLower,
		"toModelName":    toModelName,
		"toVariableName": toVariableName,
		"toBasicType":    transformer.ToBasicFieldKind,
	}
	// Define the path to the template model
	templatePath := "./orm/templates/model.tmpl"

	err := checkAndCreate(path)
	if err != nil {
		return err
	}

	fileName := extractFileNameFromModel(model.Name)
	templateName := extractFileNameFromPath(templatePath)
	tmpl, err := template.New(templateName).Funcs(functionMap).ParseFiles(templatePath)
	if err != nil {
		return err
	}

	file, err := os.Create(fmt.Sprintf("%s/%s.go", path, fileName))
	if err != nil {
		return err
	}
	defer file.Close()

	err = tmpl.Execute(file, model)
	if err != nil {
		return err
	}
	return nil
}

func toModelName(name string) string {
	if strings.Contains(name, "_") {
		parts := strings.Split(name, "_")
		for index, part := range parts {
			parts[index] = strings.Title(part)
		}
		return strings.Join(parts, "")
	}
	return strings.Title(name)
}

func toVariableName(name string) string {
	if strings.Contains(name, "_") {
		parts := strings.Split(name, "_")
		for index, part := range parts {
			if index == 0 {
				parts[index] = strings.ToLower(part)
			} else {
				parts[index] = strings.Title(part)
			}
		}
		return strings.Join(parts, "")
	}
	return strings.ToLower(name)
}

// pure function
func extractFileNameFromPath(path string) string {
	parts := strings.Split(path, "/")
	return parts[len(parts)-1]
}

// pure function
func extractFileNameFromModel(name string) string {
	fileName := strings.ReplaceAll(name, " ", "_")
	fileName = strings.ReplaceAll(fileName, "-", "_")
	return strings.ToLower(fileName)
}
